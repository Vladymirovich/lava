package chainproxy

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"strings"

	"github.com/btcsuite/btcd/btcec"
	"github.com/fullstorydev/grpcurl"
	"github.com/golang/protobuf/proto"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/dynamic"
	"github.com/jhump/protoreflect/grpcreflect"
	"github.com/lavanet/lava/relayer/chainproxy/rpcclient"
	"github.com/lavanet/lava/relayer/chainproxy/thirdparty"
	"github.com/lavanet/lava/relayer/lavasession"
	"github.com/lavanet/lava/relayer/parser"
	"github.com/lavanet/lava/relayer/sentry"
	"github.com/lavanet/lava/utils"
	pairingtypes "github.com/lavanet/lava/x/pairing/types"
	spectypes "github.com/lavanet/lava/x/spec/types"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	reflectionpbo "google.golang.org/grpc/reflection/grpc_reflection_v1alpha"
	"google.golang.org/grpc/status"
)

type Server struct {
	cp      *GrpcChainProxy
	privKey *btcec.PrivateKey
}

type GrpcMessage struct {
	methodDesc *desc.MethodDescriptor
	formatter  grpcurl.Formatter

	cp             *GrpcChainProxy
	serviceApi     *spectypes.ServiceApi
	path           string
	msg            interface{}
	requestedBlock int64
	connectionType string
	Result         json.RawMessage
}

type GrpcChainProxy struct {
	nodeUrl    string
	sentry     *sentry.Sentry
	csm        *lavasession.ConsumerSessionManager
	portalLogs *PortalLogs
	chainID    string
}

func (r *GrpcMessage) GetMsg() interface{} {
	return r.msg
}

func NewGrpcChainProxy(nodeUrl string, sentry *sentry.Sentry, csm *lavasession.ConsumerSessionManager, pLogs *PortalLogs, chainID string) ChainProxy {
	nodeUrl = strings.TrimSuffix(nodeUrl, "/")
	return &GrpcChainProxy{
		nodeUrl:    nodeUrl,
		sentry:     sentry,
		csm:        csm,
		portalLogs: pLogs,
		chainID:    chainID,
	}
}

func (m GrpcMessage) GetParams() interface{} {
	return m.msg
}

func (m GrpcMessage) GetResult() json.RawMessage { // Todo return errors here as well
	msgFactory := dynamic.NewMessageFactoryWithDefaults()
	msg := msgFactory.NewMessage(m.methodDesc.GetOutputType())
	if err := proto.Unmarshal(m.Result, msg); err != nil {
		utils.LavaFormatError("failed to unmarshal GetResult", err, nil)
		return m.Result
	}

	s, err := m.formatter(msg)
	if err != nil {
		utils.LavaFormatError("m.formatter(msg)", err, nil)
		return m.Result
	}

	return []byte(s)
}

func (m GrpcMessage) ParseBlock(inp string) (int64, error) {
	return parser.ParseDefaultBlockParameter(inp)
}

func (nm *GrpcMessage) RequestedBlock() int64 {
	return nm.requestedBlock
}

func (nm *GrpcMessage) GetServiceApi() *spectypes.ServiceApi {
	return nm.serviceApi
}

func (cp *GrpcChainProxy) GetConsumerSessionManager() *lavasession.ConsumerSessionManager {
	return cp.csm
}

func (cp *GrpcChainProxy) NewMessage(path string, data []byte) (*GrpcMessage, error) {
	//
	// Check api is supported and save it in nodeMsg
	serviceApi, err := cp.getSupportedApi(path)
	if err != nil {
		return nil, utils.LavaFormatError("failed to get supported api in NewMessage", err, &map[string]string{"path": path})
	}

	nodeMsg := &GrpcMessage{
		cp:         cp,
		serviceApi: serviceApi,
		path:       path,
		msg:        data,
	}

	return nodeMsg, nil
}

func (cp *GrpcChainProxy) FetchBlockHashByNum(ctx context.Context, blockNum int64) (string, error) {
	serviceApi, ok := cp.GetSentry().GetSpecApiByTag(spectypes.GET_BLOCK_BY_NUM)
	if !ok {
		return "", errors.New(spectypes.GET_BLOCKNUM + " tag function not found")
	}

	var nodeMsg NodeMessage
	var err error
	if serviceApi.GetParsing().FunctionTemplate != "" {
		nodeMsg, err = cp.ParseMsg(serviceApi.Name, []byte(fmt.Sprintf(serviceApi.GetParsing().FunctionTemplate, blockNum)), "")
	} else {
		nodeMsg, err = cp.NewMessage(serviceApi.Name, nil)
	}

	if err != nil {
		return "", err
	}

	_, _, _, err = nodeMsg.Send(ctx, nil)
	if err != nil {
		return "", err
	}

	blockData, err := parser.ParseMessageResponse((nodeMsg.(*GrpcMessage)), serviceApi.Parsing.ResultParsing)
	if err != nil {
		return "", err
	}

	// blockData is an interface array with the parsed result in index 0.
	// we know to expect a string result for a hash.
	return blockData[spectypes.DEFAULT_PARSED_RESULT_INDEX].(string), nil
}

func (cp *GrpcChainProxy) FetchLatestBlockNum(ctx context.Context) (int64, error) {
	serviceApi, ok := cp.GetSentry().GetSpecApiByTag(spectypes.GET_BLOCKNUM)
	if !ok {
		return spectypes.NOT_APPLICABLE, errors.New(spectypes.GET_BLOCKNUM + " tag function not found")
	}

	params := make(json.RawMessage, 0)
	nodeMsg, err := cp.NewMessage(serviceApi.GetName(), params)
	if err != nil {
		return spectypes.NOT_APPLICABLE, utils.LavaFormatError("new Message creation Failed at FetchLatestBlockNum", err, nil)
	}

	_, _, _, err = nodeMsg.Send(ctx, nil)
	if err != nil {
		return spectypes.NOT_APPLICABLE, utils.LavaFormatError("Message send Failed at FetchLatestBlockNum", err, nil)
	}

	blocknum, err := parser.ParseBlockFromReply(nodeMsg, serviceApi.Parsing.ResultParsing)
	if err != nil {
		return spectypes.NOT_APPLICABLE, err
	}

	return blocknum, nil
}

func (cp *GrpcChainProxy) GetSentry() *sentry.Sentry {
	return cp.sentry
}

func (cp *GrpcChainProxy) Start(context.Context) error {
	return nil
}

func (cp *GrpcChainProxy) getSupportedApi(path string) (*spectypes.ServiceApi, error) {
	if api, ok := cp.sentry.MatchSpecApiByName(path); ok {
		if !api.Enabled {
			return nil, fmt.Errorf("gRPC Api is disabled %s ", path)
		}
		return &api, nil
	}
	return nil, fmt.Errorf("gRPC Api not supported %s ", path)
}

func (cp *GrpcChainProxy) ParseMsg(path string, data []byte, connectionType string) (NodeMessage, error) {
	// Check API is supported and save it in nodeMsg.
	serviceApi, err := cp.getSupportedApi(path)
	if err != nil {
		return nil, utils.LavaFormatError("failed to getSupportedApi gRPC", err, nil)
	}

	nodeMsg := &GrpcMessage{
		cp:             cp,
		serviceApi:     serviceApi,
		path:           path,
		msg:            data,
		connectionType: connectionType,
	}

	return nodeMsg, nil
}

func (cp *GrpcChainProxy) PortalStart(ctx context.Context, privKey *btcec.PrivateKey, listenAddr string) {
	utils.LavaFormatInfo("gRPC PortalStart", nil)

	lis, err := net.Listen("tcp", listenAddr)
	if err != nil {
		utils.LavaFormatFatal("provider failure setting up listener", err, &map[string]string{"listenAddr": listenAddr})
	}

	sendRelayCallback := func(ctx context.Context, method string, reqBody []byte) ([]byte, error) {
		utils.LavaFormatInfo("GRPC Got Relay: "+method, nil)
		var relayReply *pairingtypes.RelayReply
		if relayReply, _, _, err = SendRelay(ctx, cp, privKey, method, string(reqBody), ""); err != nil {
			return nil, utils.LavaFormatError("Failed to SendRelay", err, nil)
		}
		return relayReply.Data, nil
	}

	s, err := thirdparty.RegisterServer(cp.chainID, sendRelayCallback)
	if err != nil {
		utils.LavaFormatFatal("provider failure RegisterServer", err, &map[string]string{"listenAddr": listenAddr})
	}

	utils.LavaFormatInfo("Server listening", &map[string]string{"Address": lis.Addr().String()})
	if err = s.Serve(lis); err != nil {
		utils.LavaFormatFatal("portal failed to serve", err, &map[string]string{"Address": lis.Addr().String()})
	}
}

func descriptorSourceFromServer(refClient *grpcreflect.Client) DescriptorSource {
	return ServerSource{Client: refClient}
}

func (nm *GrpcMessage) Send(ctx context.Context, ch chan interface{}) (relayReply *pairingtypes.RelayReply, subscriptionID string, relayReplyServer *rpcclient.ClientSubscription, err error) {
	if ch != nil {
		return nil, "", nil, utils.LavaFormatError("Subscribe is not allowed on rest", nil, nil)
	}

	connectCtx, cancel := context.WithTimeout(ctx, DefaultTimeout)
	defer cancel()
	conn, err := grpc.DialContext(connectCtx, nm.cp.nodeUrl, grpc.WithInsecure(), grpc.WithBlock()) // TODO, keep an open connection similar to others
	if err != nil {
		return nil, "", nil, utils.LavaFormatError("dialing the gRPC client failed.", err, &map[string]string{"addr": nm.cp.nodeUrl})
	}
	defer conn.Close()
	cl := grpcreflect.NewClient(ctx, reflectionpbo.NewServerReflectionClient(conn))
	descriptorSource := descriptorSourceFromServer(cl)
	svc, methodName := ParseSymbol(nm.path)
	var descriptor desc.Descriptor
	if descriptor, err = descriptorSource.FindSymbol(svc); err != nil {
		return nil, "", nil, utils.LavaFormatError("descriptorSource.FindSymbol", err, &map[string]string{"addr": nm.cp.nodeUrl})
	}

	serviceDescriptor, ok := descriptor.(*desc.ServiceDescriptor)
	if !ok {
		return nil, "", nil, utils.LavaFormatError("serviceDescriptor, ok := descriptor.(*desc.ServiceDescriptor)", err, &map[string]string{"addr": nm.cp.nodeUrl})
	}
	methodDescriptor := serviceDescriptor.FindMethodByName(methodName)
	nm.methodDesc = methodDescriptor
	msgFactory := dynamic.NewMessageFactoryWithDefaults()

	var reader io.Reader
	msg := msgFactory.NewMessage(methodDescriptor.GetInputType())
	formatMessage := false
	switch v := nm.msg.(type) {
	case []byte:
		if len(v) > 0 {
			reader = bytes.NewReader(v)
			formatMessage = true
		}
		// TODO: Test this json unmarshal should work.
		msg2 := msgFactory.NewMessage(methodDescriptor.GetInputType())
		err = json.Unmarshal(v, msg2)
		if err != nil {
			utils.LavaFormatError("Failed to json unmarshal", nil, &map[string]string{"type": fmt.Sprintf("%s", v)})
		}
	case string:
		if v != "" {
			reader = strings.NewReader(v)
			formatMessage = true
		}
	case proto.Message:
		msg = v
	default:
		return nil, "", nil, utils.LavaFormatError("Unsupported type for gRPC msg", nil, &map[string]string{"type": fmt.Sprintf("%T", v)})
	}

	rp, formatter, err := grpcurl.RequestParserAndFormatter(grpcurl.FormatJSON, descriptorSource, reader, grpcurl.FormatOptions{
		EmitJSONDefaultFields: false,
		IncludeTextSeparator:  false,
		AllowUnknownFields:    true,
	})
	if err != nil {
		return nil, "", nil, utils.LavaFormatError("Failed to create formatter", err, &map[string]string{"addr": nm.cp.nodeUrl})
	}
	nm.formatter = formatter
	if formatMessage {
		err = rp.Next(msg)
		if err != nil {
			return nil, "", nil, utils.LavaFormatError("rp.Next(msg) Failed", err, nil)
		}
	}

	response := msgFactory.NewMessage(methodDescriptor.GetOutputType())
	err = grpc.Invoke(connectCtx, nm.path, msg, response, conn)
	if err != nil {
		return nil, "", nil, utils.LavaFormatError("Invoke Failed", err, &map[string]string{"addr": nm.cp.nodeUrl, "Method": nm.path, "msg": fmt.Sprintf("%s", nm.msg)})
	}

	var respBytes []byte
	respBytes, err = proto.Marshal(response)
	if err != nil {
		return nil, "", nil, utils.LavaFormatError("proto.Marshal(response) Failed", err, &map[string]string{"addr": nm.cp.nodeUrl})
	}

	nm.Result = respBytes
	reply := &pairingtypes.RelayReply{
		Data: respBytes,
	}

	return reply, "", nil, nil
}

type DescriptorSource interface {
	// ListServices returns a list of fully-qualified service names. It will be all services in a set of
	// descriptor files or the set of all services exposed by a gRPC server.
	ListServices() ([]string, error)
	// FindSymbol returns a descriptor for the given fully-qualified symbol name.
	FindSymbol(fullyQualifiedName string) (desc.Descriptor, error)
	// AllExtensionsForType returns all known extension fields that extend the given message type name.
	AllExtensionsForType(typeName string) ([]*desc.FieldDescriptor, error)
}

type ServerSource struct {
	Client *grpcreflect.Client
}

func (ss ServerSource) ListServices() ([]string, error) {
	svcs, err := ss.Client.ListServices()
	return svcs, ReflectionSupport(err)
}

func (ss ServerSource) FindSymbol(fullyQualifiedName string) (desc.Descriptor, error) {
	file, err := ss.Client.FileContainingSymbol(fullyQualifiedName)
	if err != nil {
		return nil, ReflectionSupport(err)
	}
	d := file.FindSymbol(fullyQualifiedName)
	if d == nil {
		return nil, utils.LavaFormatError("Symbol not found", fmt.Errorf("missing symbol: %s", fullyQualifiedName), nil)
	}
	return d, nil
}

func (ss ServerSource) AllExtensionsForType(typeName string) ([]*desc.FieldDescriptor, error) {
	var exts []*desc.FieldDescriptor
	nums, err := ss.Client.AllExtensionNumbersForType(typeName)
	if err != nil {
		return nil, ReflectionSupport(err)
	}
	for _, fieldNum := range nums {
		ext, err := ss.Client.ResolveExtension(typeName, fieldNum)
		if err != nil {
			return nil, ReflectionSupport(err)
		}
		exts = append(exts, ext)
	}
	return exts, nil
}

func ReflectionSupport(err error) error {
	if err == nil {
		return nil
	}
	if stat, ok := status.FromError(err); ok && stat.Code() == codes.Unimplemented {
		return utils.LavaFormatError("server does not support the reflection API", err, nil)
	}
	return err
}

func ParseSymbol(svcAndMethod string) (string, string) {
	pos := strings.LastIndex(svcAndMethod, "/")
	if pos < 0 {
		pos = strings.LastIndex(svcAndMethod, ".")
		if pos < 0 {
			return "", ""
		}
	}
	return svcAndMethod[:pos], svcAndMethod[pos+1:]
}