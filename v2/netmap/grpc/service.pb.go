// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: v2/netmap/grpc/service.proto

package netmap

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc1 "github.com/nspcc-dev/neofs-api-go/v2/refs/grpc"
	grpc "github.com/nspcc-dev/neofs-api-go/v2/session/grpc"
	grpc2 "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Get NodeInfo structure from the particular node directly
type LocalNodeInfoRequest struct {
	// Body of the LocalNodeInfo request message
	Body *LocalNodeInfoRequest_Body `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	// Carries request meta information. Header data is used only to regulate
	// message transport and does not affect request execution.
	MetaHeader *grpc.RequestMetaHeader `protobuf:"bytes,2,opt,name=meta_header,json=metaHeader,proto3" json:"meta_header,omitempty"`
	// Carries request verification information. This header is used to
	// authenticate the nodes of the message route and check the correctness of
	// transmission.
	VerifyHeader         *grpc.RequestVerificationHeader `protobuf:"bytes,3,opt,name=verify_header,json=verifyHeader,proto3" json:"verify_header,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *LocalNodeInfoRequest) Reset()         { *m = LocalNodeInfoRequest{} }
func (m *LocalNodeInfoRequest) String() string { return proto.CompactTextString(m) }
func (*LocalNodeInfoRequest) ProtoMessage()    {}
func (*LocalNodeInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8a816fa1f32c636, []int{0}
}
func (m *LocalNodeInfoRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LocalNodeInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LocalNodeInfoRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LocalNodeInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocalNodeInfoRequest.Merge(m, src)
}
func (m *LocalNodeInfoRequest) XXX_Size() int {
	return m.Size()
}
func (m *LocalNodeInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LocalNodeInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LocalNodeInfoRequest proto.InternalMessageInfo

func (m *LocalNodeInfoRequest) GetBody() *LocalNodeInfoRequest_Body {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *LocalNodeInfoRequest) GetMetaHeader() *grpc.RequestMetaHeader {
	if m != nil {
		return m.MetaHeader
	}
	return nil
}

func (m *LocalNodeInfoRequest) GetVerifyHeader() *grpc.RequestVerificationHeader {
	if m != nil {
		return m.VerifyHeader
	}
	return nil
}

// LocalNodeInfo request body is empty.
type LocalNodeInfoRequest_Body struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LocalNodeInfoRequest_Body) Reset()         { *m = LocalNodeInfoRequest_Body{} }
func (m *LocalNodeInfoRequest_Body) String() string { return proto.CompactTextString(m) }
func (*LocalNodeInfoRequest_Body) ProtoMessage()    {}
func (*LocalNodeInfoRequest_Body) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8a816fa1f32c636, []int{0, 0}
}
func (m *LocalNodeInfoRequest_Body) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LocalNodeInfoRequest_Body) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LocalNodeInfoRequest_Body.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LocalNodeInfoRequest_Body) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocalNodeInfoRequest_Body.Merge(m, src)
}
func (m *LocalNodeInfoRequest_Body) XXX_Size() int {
	return m.Size()
}
func (m *LocalNodeInfoRequest_Body) XXX_DiscardUnknown() {
	xxx_messageInfo_LocalNodeInfoRequest_Body.DiscardUnknown(m)
}

var xxx_messageInfo_LocalNodeInfoRequest_Body proto.InternalMessageInfo

// Local Node Info, including API Version in use
type LocalNodeInfoResponse struct {
	// Body of the balance response message.
	Body *LocalNodeInfoResponse_Body `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	// Carries response meta information. Header data is used only to regulate
	// message transport and does not affect response execution.
	MetaHeader *grpc.ResponseMetaHeader `protobuf:"bytes,2,opt,name=meta_header,json=metaHeader,proto3" json:"meta_header,omitempty"`
	// Carries response verification information. This header is used to
	// authenticate the nodes of the message route and check the correctness of
	// transmission.
	VerifyHeader         *grpc.ResponseVerificationHeader `protobuf:"bytes,3,opt,name=verify_header,json=verifyHeader,proto3" json:"verify_header,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *LocalNodeInfoResponse) Reset()         { *m = LocalNodeInfoResponse{} }
func (m *LocalNodeInfoResponse) String() string { return proto.CompactTextString(m) }
func (*LocalNodeInfoResponse) ProtoMessage()    {}
func (*LocalNodeInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8a816fa1f32c636, []int{1}
}
func (m *LocalNodeInfoResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LocalNodeInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LocalNodeInfoResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LocalNodeInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocalNodeInfoResponse.Merge(m, src)
}
func (m *LocalNodeInfoResponse) XXX_Size() int {
	return m.Size()
}
func (m *LocalNodeInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LocalNodeInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LocalNodeInfoResponse proto.InternalMessageInfo

func (m *LocalNodeInfoResponse) GetBody() *LocalNodeInfoResponse_Body {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *LocalNodeInfoResponse) GetMetaHeader() *grpc.ResponseMetaHeader {
	if m != nil {
		return m.MetaHeader
	}
	return nil
}

func (m *LocalNodeInfoResponse) GetVerifyHeader() *grpc.ResponseVerificationHeader {
	if m != nil {
		return m.VerifyHeader
	}
	return nil
}

// Local Node Info, including API Version in use.
type LocalNodeInfoResponse_Body struct {
	// Latest NeoFS API version in use
	Version *grpc1.Version `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	// NodeInfo structure with recent information from node itself
	NodeInfo             *NodeInfo `protobuf:"bytes,2,opt,name=node_info,json=nodeInfo,proto3" json:"node_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *LocalNodeInfoResponse_Body) Reset()         { *m = LocalNodeInfoResponse_Body{} }
func (m *LocalNodeInfoResponse_Body) String() string { return proto.CompactTextString(m) }
func (*LocalNodeInfoResponse_Body) ProtoMessage()    {}
func (*LocalNodeInfoResponse_Body) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8a816fa1f32c636, []int{1, 0}
}
func (m *LocalNodeInfoResponse_Body) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LocalNodeInfoResponse_Body) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LocalNodeInfoResponse_Body.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LocalNodeInfoResponse_Body) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocalNodeInfoResponse_Body.Merge(m, src)
}
func (m *LocalNodeInfoResponse_Body) XXX_Size() int {
	return m.Size()
}
func (m *LocalNodeInfoResponse_Body) XXX_DiscardUnknown() {
	xxx_messageInfo_LocalNodeInfoResponse_Body.DiscardUnknown(m)
}

var xxx_messageInfo_LocalNodeInfoResponse_Body proto.InternalMessageInfo

func (m *LocalNodeInfoResponse_Body) GetVersion() *grpc1.Version {
	if m != nil {
		return m.Version
	}
	return nil
}

func (m *LocalNodeInfoResponse_Body) GetNodeInfo() *NodeInfo {
	if m != nil {
		return m.NodeInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*LocalNodeInfoRequest)(nil), "neo.fs.v2.netmap.LocalNodeInfoRequest")
	proto.RegisterType((*LocalNodeInfoRequest_Body)(nil), "neo.fs.v2.netmap.LocalNodeInfoRequest.Body")
	proto.RegisterType((*LocalNodeInfoResponse)(nil), "neo.fs.v2.netmap.LocalNodeInfoResponse")
	proto.RegisterType((*LocalNodeInfoResponse_Body)(nil), "neo.fs.v2.netmap.LocalNodeInfoResponse.Body")
}

func init() { proto.RegisterFile("v2/netmap/grpc/service.proto", fileDescriptor_d8a816fa1f32c636) }

var fileDescriptor_d8a816fa1f32c636 = []byte{
	// 438 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xdf, 0x6a, 0xd4, 0x40,
	0x14, 0xc6, 0xcd, 0x5a, 0xaa, 0x4e, 0x5d, 0x90, 0x51, 0x71, 0x89, 0x12, 0xa4, 0xf8, 0x0f, 0x6c,
	0x66, 0x30, 0x5e, 0xf4, 0xc2, 0x0b, 0xb5, 0x60, 0xb1, 0xa0, 0x8b, 0xa6, 0xd0, 0x0b, 0x6f, 0xd6,
	0x6c, 0x72, 0xb2, 0x1d, 0x30, 0x73, 0xd2, 0x99, 0xe9, 0x40, 0xde, 0xc4, 0x67, 0xf0, 0x49, 0xbc,
	0xf4, 0x11, 0x64, 0xbd, 0xf2, 0x11, 0xbc, 0x93, 0x64, 0xd2, 0xb2, 0x7f, 0xd2, 0xee, 0xde, 0x9d,
	0xe5, 0x7c, 0xdf, 0xb7, 0xe7, 0x77, 0x72, 0x86, 0x3c, 0xb0, 0x11, 0x97, 0x60, 0x8a, 0xa4, 0xe4,
	0x13, 0x55, 0xa6, 0x5c, 0x83, 0xb2, 0x22, 0x05, 0x56, 0x2a, 0x34, 0x48, 0x6f, 0x49, 0x40, 0x96,
	0x6b, 0x66, 0x23, 0xe6, 0x44, 0xbe, 0xbf, 0xa0, 0x37, 0x55, 0x09, 0xda, 0xa9, 0xfd, 0x81, 0x8d,
	0xb8, 0x82, 0x5c, 0x2f, 0x77, 0xee, 0xdb, 0x88, 0x6b, 0xd0, 0x5a, 0xa0, 0x5c, 0x6a, 0x6e, 0xff,
	0xf3, 0xc8, 0x9d, 0x0f, 0x98, 0x26, 0xdf, 0x86, 0x98, 0xc1, 0x81, 0xcc, 0x31, 0x86, 0x93, 0x53,
	0xd0, 0x86, 0xbe, 0x26, 0x1b, 0x63, 0xcc, 0xaa, 0x81, 0xf7, 0xd0, 0x7b, 0xb6, 0x15, 0x3d, 0x67,
	0x8b, 0xc3, 0xb0, 0x2e, 0x17, 0xdb, 0xc3, 0xac, 0x8a, 0x1b, 0x23, 0x7d, 0x47, 0xb6, 0x0a, 0x30,
	0xc9, 0xe8, 0x18, 0x92, 0x0c, 0xd4, 0xa0, 0xd7, 0xe4, 0x3c, 0x9a, 0xc9, 0x69, 0x67, 0x62, 0xad,
	0xf7, 0x23, 0x98, 0xe4, 0x7d, 0xa3, 0x8d, 0x49, 0x71, 0x5e, 0xd3, 0xcf, 0xa4, 0x6f, 0x41, 0x89,
	0xbc, 0x3a, 0x0b, 0xba, 0xda, 0x04, 0xed, 0x5c, 0x1c, 0x74, 0x54, 0xcb, 0x45, 0x9a, 0x18, 0x81,
	0xb2, 0x0d, 0xbc, 0xe9, 0x22, 0xdc, 0x2f, 0x7f, 0x93, 0x6c, 0xd4, 0x73, 0x6e, 0xff, 0xed, 0x91,
	0xbb, 0x0b, 0x14, 0xba, 0x44, 0xa9, 0x81, 0xbe, 0x99, 0x83, 0xdf, 0x59, 0x09, 0xef, 0x6c, 0xb3,
	0xf4, 0xfb, 0x5d, 0xf4, 0x8f, 0x3b, 0x87, 0x76, 0xe6, 0x0b, 0xf0, 0xe3, 0x6e, 0xfc, 0xf0, 0x92,
	0xa4, 0x95, 0xfc, 0xca, 0xf1, 0xd3, 0x17, 0xe4, 0x9a, 0x05, 0x55, 0x7b, 0x5b, 0xd0, 0x7b, 0x33,
	0xa9, 0xf5, 0x2d, 0xb1, 0x23, 0xd7, 0x8e, 0xcf, 0x74, 0x74, 0x97, 0xdc, 0x90, 0x98, 0xc1, 0x48,
	0xc8, 0x1c, 0x5b, 0x28, 0x7f, 0x79, 0x3b, 0xe7, 0x8b, 0xb9, 0x2e, 0xdb, 0x2a, 0x3a, 0x21, 0xfd,
	0x61, 0xd3, 0x3c, 0x74, 0x37, 0x4e, 0xbf, 0x92, 0xfe, 0xdc, 0x12, 0xe9, 0x93, 0xf5, 0x4e, 0xcc,
	0x7f, 0xba, 0xe6, 0xd7, 0xd8, 0x1b, 0xfd, 0x9c, 0x06, 0xde, 0xaf, 0x69, 0xe0, 0xfd, 0x9e, 0x06,
	0xde, 0xf7, 0x3f, 0xc1, 0x95, 0x2f, 0xbb, 0x13, 0x61, 0x8e, 0x4f, 0xc7, 0x2c, 0xc5, 0x82, 0x4b,
	0x5d, 0xa6, 0x69, 0x98, 0x81, 0xe5, 0x12, 0x30, 0xd7, 0x61, 0x52, 0x8a, 0x70, 0x82, 0x7c, 0xfe,
	0x85, 0xbd, 0x72, 0xf5, 0x8f, 0xde, 0xed, 0x21, 0xe0, 0xfe, 0x21, 0x7b, 0xfb, 0xe9, 0xa0, 0xfe,
	0x5b, 0x47, 0x32, 0xde, 0x6c, 0x9e, 0xd0, 0xcb, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0xad, 0xd6,
	0xfb, 0x46, 0xc7, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc2.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc2.SupportPackageIsVersion4

// NetmapServiceClient is the client API for NetmapService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NetmapServiceClient interface {
	// Get NodeInfo structure from the particular node directly. Node information
	// can be taken from `Netmap` smart contract, but in some cases the one may
	// want to get recent information directly, or to talk to the node not yet
	// present in `Network Map` to find out what API version can be used for
	// further communication. Can also be used to check if node is up and running.
	LocalNodeInfo(ctx context.Context, in *LocalNodeInfoRequest, opts ...grpc2.CallOption) (*LocalNodeInfoResponse, error)
}

type netmapServiceClient struct {
	cc *grpc2.ClientConn
}

func NewNetmapServiceClient(cc *grpc2.ClientConn) NetmapServiceClient {
	return &netmapServiceClient{cc}
}

func (c *netmapServiceClient) LocalNodeInfo(ctx context.Context, in *LocalNodeInfoRequest, opts ...grpc2.CallOption) (*LocalNodeInfoResponse, error) {
	out := new(LocalNodeInfoResponse)
	err := c.cc.Invoke(ctx, "/neo.fs.v2.netmap.NetmapService/LocalNodeInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NetmapServiceServer is the server API for NetmapService service.
type NetmapServiceServer interface {
	// Get NodeInfo structure from the particular node directly. Node information
	// can be taken from `Netmap` smart contract, but in some cases the one may
	// want to get recent information directly, or to talk to the node not yet
	// present in `Network Map` to find out what API version can be used for
	// further communication. Can also be used to check if node is up and running.
	LocalNodeInfo(context.Context, *LocalNodeInfoRequest) (*LocalNodeInfoResponse, error)
}

// UnimplementedNetmapServiceServer can be embedded to have forward compatible implementations.
type UnimplementedNetmapServiceServer struct {
}

func (*UnimplementedNetmapServiceServer) LocalNodeInfo(ctx context.Context, req *LocalNodeInfoRequest) (*LocalNodeInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LocalNodeInfo not implemented")
}

func RegisterNetmapServiceServer(s *grpc2.Server, srv NetmapServiceServer) {
	s.RegisterService(&_NetmapService_serviceDesc, srv)
}

func _NetmapService_LocalNodeInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc2.UnaryServerInterceptor) (interface{}, error) {
	in := new(LocalNodeInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetmapServiceServer).LocalNodeInfo(ctx, in)
	}
	info := &grpc2.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/neo.fs.v2.netmap.NetmapService/LocalNodeInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetmapServiceServer).LocalNodeInfo(ctx, req.(*LocalNodeInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NetmapService_serviceDesc = grpc2.ServiceDesc{
	ServiceName: "neo.fs.v2.netmap.NetmapService",
	HandlerType: (*NetmapServiceServer)(nil),
	Methods: []grpc2.MethodDesc{
		{
			MethodName: "LocalNodeInfo",
			Handler:    _NetmapService_LocalNodeInfo_Handler,
		},
	},
	Streams:  []grpc2.StreamDesc{},
	Metadata: "v2/netmap/grpc/service.proto",
}

func (m *LocalNodeInfoRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LocalNodeInfoRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LocalNodeInfoRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.VerifyHeader != nil {
		{
			size, err := m.VerifyHeader.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintService(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.MetaHeader != nil {
		{
			size, err := m.MetaHeader.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintService(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Body != nil {
		{
			size, err := m.Body.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintService(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *LocalNodeInfoRequest_Body) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LocalNodeInfoRequest_Body) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LocalNodeInfoRequest_Body) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	return len(dAtA) - i, nil
}

func (m *LocalNodeInfoResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LocalNodeInfoResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LocalNodeInfoResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.VerifyHeader != nil {
		{
			size, err := m.VerifyHeader.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintService(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.MetaHeader != nil {
		{
			size, err := m.MetaHeader.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintService(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Body != nil {
		{
			size, err := m.Body.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintService(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *LocalNodeInfoResponse_Body) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LocalNodeInfoResponse_Body) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LocalNodeInfoResponse_Body) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.NodeInfo != nil {
		{
			size, err := m.NodeInfo.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintService(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Version != nil {
		{
			size, err := m.Version.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintService(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintService(dAtA []byte, offset int, v uint64) int {
	offset -= sovService(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *LocalNodeInfoRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Body != nil {
		l = m.Body.Size()
		n += 1 + l + sovService(uint64(l))
	}
	if m.MetaHeader != nil {
		l = m.MetaHeader.Size()
		n += 1 + l + sovService(uint64(l))
	}
	if m.VerifyHeader != nil {
		l = m.VerifyHeader.Size()
		n += 1 + l + sovService(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *LocalNodeInfoRequest_Body) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *LocalNodeInfoResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Body != nil {
		l = m.Body.Size()
		n += 1 + l + sovService(uint64(l))
	}
	if m.MetaHeader != nil {
		l = m.MetaHeader.Size()
		n += 1 + l + sovService(uint64(l))
	}
	if m.VerifyHeader != nil {
		l = m.VerifyHeader.Size()
		n += 1 + l + sovService(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *LocalNodeInfoResponse_Body) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Version != nil {
		l = m.Version.Size()
		n += 1 + l + sovService(uint64(l))
	}
	if m.NodeInfo != nil {
		l = m.NodeInfo.Size()
		n += 1 + l + sovService(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovService(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozService(x uint64) (n int) {
	return sovService(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LocalNodeInfoRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
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
			return fmt.Errorf("proto: LocalNodeInfoRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LocalNodeInfoRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Body", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Body == nil {
				m.Body = &LocalNodeInfoRequest_Body{}
			}
			if err := m.Body.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MetaHeader", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MetaHeader == nil {
				m.MetaHeader = &grpc.RequestMetaHeader{}
			}
			if err := m.MetaHeader.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VerifyHeader", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.VerifyHeader == nil {
				m.VerifyHeader = &grpc.RequestVerificationHeader{}
			}
			if err := m.VerifyHeader.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthService
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *LocalNodeInfoRequest_Body) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
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
			return fmt.Errorf("proto: Body: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Body: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthService
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *LocalNodeInfoResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
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
			return fmt.Errorf("proto: LocalNodeInfoResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LocalNodeInfoResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Body", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Body == nil {
				m.Body = &LocalNodeInfoResponse_Body{}
			}
			if err := m.Body.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MetaHeader", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MetaHeader == nil {
				m.MetaHeader = &grpc.ResponseMetaHeader{}
			}
			if err := m.MetaHeader.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VerifyHeader", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.VerifyHeader == nil {
				m.VerifyHeader = &grpc.ResponseVerificationHeader{}
			}
			if err := m.VerifyHeader.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthService
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *LocalNodeInfoResponse_Body) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
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
			return fmt.Errorf("proto: Body: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Body: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Version == nil {
				m.Version = &grpc1.Version{}
			}
			if err := m.Version.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.NodeInfo == nil {
				m.NodeInfo = &NodeInfo{}
			}
			if err := m.NodeInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthService
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipService(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowService
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
					return 0, ErrIntOverflowService
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
					return 0, ErrIntOverflowService
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
				return 0, ErrInvalidLengthService
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupService
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthService
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthService        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowService          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupService = fmt.Errorf("proto: unexpected end of group")
)
