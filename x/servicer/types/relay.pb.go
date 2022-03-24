// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: servicer/relay.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type RelayRequest struct {
	SpecId      uint32 `protobuf:"varint,1,opt,name=spec_id,json=specId,proto3" json:"spec_id,omitempty"`
	ApiId       uint32 `protobuf:"varint,2,opt,name=api_id,json=apiId,proto3" json:"api_id,omitempty"`
	SessionId   uint64 `protobuf:"varint,3,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	CuSum       uint64 `protobuf:"varint,4,opt,name=cu_sum,json=cuSum,proto3" json:"cu_sum,omitempty"`
	Data        []byte `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
	Sig         []byte `protobuf:"bytes,6,opt,name=sig,proto3" json:"sig,omitempty"`
	Servicer    string `protobuf:"bytes,7,opt,name=servicer,proto3" json:"servicer,omitempty"`
	BlockHeight int64  `protobuf:"varint,8,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
}

func (m *RelayRequest) Reset()         { *m = RelayRequest{} }
func (m *RelayRequest) String() string { return proto.CompactTextString(m) }
func (*RelayRequest) ProtoMessage()    {}
func (*RelayRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_997f5cfcc544102b, []int{0}
}
func (m *RelayRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RelayRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RelayRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RelayRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RelayRequest.Merge(m, src)
}
func (m *RelayRequest) XXX_Size() int {
	return m.Size()
}
func (m *RelayRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RelayRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RelayRequest proto.InternalMessageInfo

func (m *RelayRequest) GetSpecId() uint32 {
	if m != nil {
		return m.SpecId
	}
	return 0
}

func (m *RelayRequest) GetApiId() uint32 {
	if m != nil {
		return m.ApiId
	}
	return 0
}

func (m *RelayRequest) GetSessionId() uint64 {
	if m != nil {
		return m.SessionId
	}
	return 0
}

func (m *RelayRequest) GetCuSum() uint64 {
	if m != nil {
		return m.CuSum
	}
	return 0
}

func (m *RelayRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *RelayRequest) GetSig() []byte {
	if m != nil {
		return m.Sig
	}
	return nil
}

func (m *RelayRequest) GetServicer() string {
	if m != nil {
		return m.Servicer
	}
	return ""
}

func (m *RelayRequest) GetBlockHeight() int64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

type RelayReply struct {
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Sig  []byte `protobuf:"bytes,2,opt,name=sig,proto3" json:"sig,omitempty"`
}

func (m *RelayReply) Reset()         { *m = RelayReply{} }
func (m *RelayReply) String() string { return proto.CompactTextString(m) }
func (*RelayReply) ProtoMessage()    {}
func (*RelayReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_997f5cfcc544102b, []int{1}
}
func (m *RelayReply) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RelayReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RelayReply.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RelayReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RelayReply.Merge(m, src)
}
func (m *RelayReply) XXX_Size() int {
	return m.Size()
}
func (m *RelayReply) XXX_DiscardUnknown() {
	xxx_messageInfo_RelayReply.DiscardUnknown(m)
}

var xxx_messageInfo_RelayReply proto.InternalMessageInfo

func (m *RelayReply) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *RelayReply) GetSig() []byte {
	if m != nil {
		return m.Sig
	}
	return nil
}

func init() {
	proto.RegisterType((*RelayRequest)(nil), "lavanet.lava.servicer.RelayRequest")
	proto.RegisterType((*RelayReply)(nil), "lavanet.lava.servicer.RelayReply")
}

func init() { proto.RegisterFile("servicer/relay.proto", fileDescriptor_997f5cfcc544102b) }

var fileDescriptor_997f5cfcc544102b = []byte{
	// 328 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xc1, 0x4e, 0xf2, 0x40,
	0x10, 0xc7, 0xbb, 0x40, 0x0b, 0xcc, 0xc7, 0x97, 0x98, 0x8d, 0xc4, 0x86, 0xc4, 0xa6, 0xe0, 0xa5,
	0xa7, 0x36, 0xc1, 0x37, 0xe0, 0x24, 0x47, 0xd7, 0x9b, 0x31, 0x21, 0x4b, 0xbb, 0x81, 0x8d, 0x85,
	0xae, 0xdd, 0x5d, 0x62, 0xdf, 0xc2, 0xc7, 0xf2, 0xc8, 0x51, 0x6f, 0x06, 0x5e, 0xc4, 0xec, 0x5a,
	0x1a, 0x13, 0x8d, 0xa7, 0x99, 0xf9, 0xcd, 0x7f, 0xfa, 0xef, 0xce, 0xc0, 0xb9, 0x64, 0xe5, 0x8e,
	0xa7, 0xac, 0x4c, 0x4a, 0x96, 0xd3, 0x2a, 0x16, 0x65, 0xa1, 0x0a, 0x3c, 0xcc, 0xe9, 0x8e, 0x6e,
	0x99, 0x8a, 0x4d, 0x8c, 0x4f, 0x92, 0xc9, 0x3b, 0x82, 0x01, 0x31, 0x32, 0xc2, 0x9e, 0x34, 0x93,
	0x0a, 0x5f, 0x40, 0x57, 0x0a, 0x96, 0x2e, 0x78, 0xe6, 0xa3, 0x10, 0x45, 0xff, 0x89, 0x67, 0xca,
	0x79, 0x86, 0x87, 0xe0, 0x51, 0xc1, 0x0d, 0x6f, 0x59, 0xee, 0x52, 0xc1, 0xe7, 0x19, 0xbe, 0x04,
	0x90, 0x4c, 0x4a, 0x5e, 0x6c, 0x4d, 0xab, 0x1d, 0xa2, 0xa8, 0x43, 0xfa, 0x35, 0xf9, 0x9a, 0x4a,
	0xf5, 0x42, 0xea, 0x8d, 0xdf, 0xb1, 0x2d, 0x37, 0xd5, 0x77, 0x7a, 0x83, 0x31, 0x74, 0x32, 0xaa,
	0xa8, 0xef, 0x86, 0x28, 0x1a, 0x10, 0x9b, 0xe3, 0x33, 0x68, 0x4b, 0xbe, 0xf2, 0x3d, 0x8b, 0x4c,
	0x8a, 0x47, 0xd0, 0x3b, 0xfd, 0xa8, 0xdf, 0x0d, 0x51, 0xd4, 0x27, 0x4d, 0x8d, 0xc7, 0x30, 0x58,
	0xe6, 0x45, 0xfa, 0xb8, 0x58, 0x33, 0xbe, 0x5a, 0x2b, 0xbf, 0x17, 0xa2, 0xa8, 0x4d, 0xfe, 0x59,
	0x76, 0x63, 0xd1, 0x64, 0x0a, 0x50, 0x3f, 0x4d, 0xe4, 0x55, 0x63, 0x89, 0x7e, 0x5a, 0xb6, 0x1a,
	0xcb, 0xe9, 0x03, 0x74, 0xed, 0x0c, 0x2b, 0xf1, 0x2d, 0xb8, 0x36, 0xc5, 0x57, 0xf1, 0xaf, 0xbb,
	0x8b, 0xbf, 0xef, 0x6d, 0x34, 0xfe, 0x5b, 0x24, 0xf2, 0x6a, 0xe2, 0xcc, 0x66, 0xaf, 0x87, 0x00,
	0xed, 0x0f, 0x01, 0xfa, 0x38, 0x04, 0xe8, 0xe5, 0x18, 0x38, 0xfb, 0x63, 0xe0, 0xbc, 0x1d, 0x03,
	0xe7, 0x3e, 0x5a, 0x71, 0xb5, 0xd6, 0xcb, 0x38, 0x2d, 0x36, 0x49, 0xfd, 0x21, 0x1b, 0x93, 0xe7,
	0xa4, 0x39, 0xa7, 0xaa, 0x04, 0x93, 0x4b, 0xcf, 0xde, 0xf3, 0xfa, 0x33, 0x00, 0x00, 0xff, 0xff,
	0x8e, 0xcf, 0x76, 0x3b, 0xe7, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RelayerClient is the client API for Relayer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RelayerClient interface {
	Relay(ctx context.Context, in *RelayRequest, opts ...grpc.CallOption) (*RelayReply, error)
}

type relayerClient struct {
	cc grpc1.ClientConn
}

func NewRelayerClient(cc grpc1.ClientConn) RelayerClient {
	return &relayerClient{cc}
}

func (c *relayerClient) Relay(ctx context.Context, in *RelayRequest, opts ...grpc.CallOption) (*RelayReply, error) {
	out := new(RelayReply)
	err := c.cc.Invoke(ctx, "/lavanet.lava.servicer.Relayer/Relay", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RelayerServer is the server API for Relayer service.
type RelayerServer interface {
	Relay(context.Context, *RelayRequest) (*RelayReply, error)
}

// UnimplementedRelayerServer can be embedded to have forward compatible implementations.
type UnimplementedRelayerServer struct {
}

func (*UnimplementedRelayerServer) Relay(ctx context.Context, req *RelayRequest) (*RelayReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Relay not implemented")
}

func RegisterRelayerServer(s grpc1.Server, srv RelayerServer) {
	s.RegisterService(&_Relayer_serviceDesc, srv)
}

func _Relayer_Relay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RelayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelayerServer).Relay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lavanet.lava.servicer.Relayer/Relay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelayerServer).Relay(ctx, req.(*RelayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Relayer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "lavanet.lava.servicer.Relayer",
	HandlerType: (*RelayerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Relay",
			Handler:    _Relayer_Relay_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "servicer/relay.proto",
}

func (m *RelayRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RelayRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RelayRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BlockHeight != 0 {
		i = encodeVarintRelay(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x40
	}
	if len(m.Servicer) > 0 {
		i -= len(m.Servicer)
		copy(dAtA[i:], m.Servicer)
		i = encodeVarintRelay(dAtA, i, uint64(len(m.Servicer)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Sig) > 0 {
		i -= len(m.Sig)
		copy(dAtA[i:], m.Sig)
		i = encodeVarintRelay(dAtA, i, uint64(len(m.Sig)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintRelay(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x2a
	}
	if m.CuSum != 0 {
		i = encodeVarintRelay(dAtA, i, uint64(m.CuSum))
		i--
		dAtA[i] = 0x20
	}
	if m.SessionId != 0 {
		i = encodeVarintRelay(dAtA, i, uint64(m.SessionId))
		i--
		dAtA[i] = 0x18
	}
	if m.ApiId != 0 {
		i = encodeVarintRelay(dAtA, i, uint64(m.ApiId))
		i--
		dAtA[i] = 0x10
	}
	if m.SpecId != 0 {
		i = encodeVarintRelay(dAtA, i, uint64(m.SpecId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *RelayReply) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RelayReply) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RelayReply) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Sig) > 0 {
		i -= len(m.Sig)
		copy(dAtA[i:], m.Sig)
		i = encodeVarintRelay(dAtA, i, uint64(len(m.Sig)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintRelay(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintRelay(dAtA []byte, offset int, v uint64) int {
	offset -= sovRelay(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RelayRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SpecId != 0 {
		n += 1 + sovRelay(uint64(m.SpecId))
	}
	if m.ApiId != 0 {
		n += 1 + sovRelay(uint64(m.ApiId))
	}
	if m.SessionId != 0 {
		n += 1 + sovRelay(uint64(m.SessionId))
	}
	if m.CuSum != 0 {
		n += 1 + sovRelay(uint64(m.CuSum))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovRelay(uint64(l))
	}
	l = len(m.Sig)
	if l > 0 {
		n += 1 + l + sovRelay(uint64(l))
	}
	l = len(m.Servicer)
	if l > 0 {
		n += 1 + l + sovRelay(uint64(l))
	}
	if m.BlockHeight != 0 {
		n += 1 + sovRelay(uint64(m.BlockHeight))
	}
	return n
}

func (m *RelayReply) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovRelay(uint64(l))
	}
	l = len(m.Sig)
	if l > 0 {
		n += 1 + l + sovRelay(uint64(l))
	}
	return n
}

func sovRelay(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRelay(x uint64) (n int) {
	return sovRelay(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RelayRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRelay
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RelayRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RelayRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SpecId", wireType)
			}
			m.SpecId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRelay
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SpecId |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApiId", wireType)
			}
			m.ApiId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRelay
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ApiId |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SessionId", wireType)
			}
			m.SessionId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRelay
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SessionId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CuSum", wireType)
			}
			m.CuSum = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRelay
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CuSum |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRelay
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthRelay
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthRelay
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sig", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRelay
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthRelay
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthRelay
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sig = append(m.Sig[:0], dAtA[iNdEx:postIndex]...)
			if m.Sig == nil {
				m.Sig = []byte{}
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Servicer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRelay
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRelay
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRelay
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Servicer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRelay
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRelay(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRelay
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RelayReply) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRelay
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RelayReply: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RelayReply: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRelay
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthRelay
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthRelay
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sig", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRelay
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthRelay
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthRelay
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sig = append(m.Sig[:0], dAtA[iNdEx:postIndex]...)
			if m.Sig == nil {
				m.Sig = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRelay(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRelay
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipRelay(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRelay
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRelay
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRelay
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthRelay
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRelay
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRelay
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRelay        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRelay          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRelay = fmt.Errorf("proto: unexpected end of group")
)
