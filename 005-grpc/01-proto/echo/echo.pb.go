// Code generated test by protoc-gen-go. DO NOT EDIT.
// source: echo.proto

package echo

import (
	context "context"
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type EchoRequest struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EchoRequest) Reset()         { *m = EchoRequest{} }
func (m *EchoRequest) String() string { return proto.CompactTextString(m) }
func (*EchoRequest) ProtoMessage()    {}
func (*EchoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_08134aea513e0001, []int{0}
}

func (m *EchoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EchoRequest.Unmarshal(m, b)
}
func (m *EchoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EchoRequest.Marshal(b, m, deterministic)
}
func (m *EchoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EchoRequest.Merge(m, src)
}
func (m *EchoRequest) XXX_Size() int {
	return xxx_messageInfo_EchoRequest.Size(m)
}
func (m *EchoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EchoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EchoRequest proto.InternalMessageInfo

func (m *EchoRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type EchoResponse struct {
	Response             string   `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EchoResponse) Reset()         { *m = EchoResponse{} }
func (m *EchoResponse) String() string { return proto.CompactTextString(m) }
func (*EchoResponse) ProtoMessage()    {}
func (*EchoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_08134aea513e0001, []int{1}
}

func (m *EchoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EchoResponse.Unmarshal(m, b)
}
func (m *EchoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EchoResponse.Marshal(b, m, deterministic)
}
func (m *EchoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EchoResponse.Merge(m, src)
}
func (m *EchoResponse) XXX_Size() int {
	return xxx_messageInfo_EchoResponse.Size(m)
}
func (m *EchoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EchoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EchoResponse proto.InternalMessageInfo

func (m *EchoResponse) GetResponse() string {
	if m != nil {
		return m.Response
	}
	return ""
}

func init() {
	proto.RegisterType((*EchoRequest)(nil), "EchoRequest")
	proto.RegisterType((*EchoResponse)(nil), "EchoResponse")
}

func init() { proto.RegisterFile("echo.proto", fileDescriptor_08134aea513e0001) }

var fileDescriptor_08134aea513e0001 = []byte{
	// 129 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x4d, 0xce, 0xc8,
	0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x52, 0xe7, 0xe2, 0x76, 0x4d, 0xce, 0xc8, 0x0f, 0x4a,
	0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x92, 0xe0, 0x62, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f,
	0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x71, 0x95, 0xb4, 0xb8, 0x78, 0x20, 0x0a, 0x8b,
	0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0xa4, 0xb8, 0x38, 0x8a, 0xa0, 0x6c, 0xa8, 0x52, 0x38, 0xdf,
	0xc8, 0x98, 0x8b, 0x0b, 0xa4, 0x36, 0x38, 0xb5, 0xa8, 0x2c, 0xb5, 0x48, 0x48, 0x95, 0x8b, 0x05,
	0xc4, 0x13, 0xe2, 0xd1, 0x43, 0xb2, 0x49, 0x8a, 0x57, 0x0f, 0xd9, 0x38, 0x25, 0x86, 0x24, 0x36,
	0xb0, 0x83, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x88, 0x1c, 0x34, 0x83, 0x9e, 0x00, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EchoServerClient is the client API for EchoServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EchoServerClient interface {
	Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error)
}

type echoServerClient struct {
	cc *grpc.ClientConn
}

func NewEchoServerClient(cc *grpc.ClientConn) EchoServerClient {
	return &echoServerClient{cc}
}

func (c *echoServerClient) Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error) {
	out := new(EchoResponse)
	err := c.cc.Invoke(ctx, "/EchoServer/Echo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EchoServerServer is the server API for EchoServer service.
type EchoServerServer interface {
	Echo(context.Context, *EchoRequest) (*EchoResponse, error)
}

func RegisterEchoServerServer(s *grpc.Server, srv EchoServerServer) {
	s.RegisterService(&_EchoServer_serviceDesc, srv)
}

func _EchoServer_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EchoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServerServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/EchoServer/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServerServer).Echo(ctx, req.(*EchoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EchoServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "EchoServer",
	HandlerType: (*EchoServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _EchoServer_Echo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "echo.proto",
}
