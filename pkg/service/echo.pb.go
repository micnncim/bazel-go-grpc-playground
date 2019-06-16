// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service/echo.proto

package service

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

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
	return fileDescriptor_echo_ff2fe5810199bc2f, []int{0}
}
func (m *EchoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EchoRequest.Unmarshal(m, b)
}
func (m *EchoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EchoRequest.Marshal(b, m, deterministic)
}
func (dst *EchoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EchoRequest.Merge(dst, src)
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
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EchoResponse) Reset()         { *m = EchoResponse{} }
func (m *EchoResponse) String() string { return proto.CompactTextString(m) }
func (*EchoResponse) ProtoMessage()    {}
func (*EchoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_echo_ff2fe5810199bc2f, []int{1}
}
func (m *EchoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EchoResponse.Unmarshal(m, b)
}
func (m *EchoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EchoResponse.Marshal(b, m, deterministic)
}
func (dst *EchoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EchoResponse.Merge(dst, src)
}
func (m *EchoResponse) XXX_Size() int {
	return xxx_messageInfo_EchoResponse.Size(m)
}
func (m *EchoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EchoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EchoResponse proto.InternalMessageInfo

func (m *EchoResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*EchoRequest)(nil), "service.EchoRequest")
	proto.RegisterType((*EchoResponse)(nil), "service.EchoResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EchoAPIClient is the client API for EchoAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EchoAPIClient interface {
	Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error)
}

type echoAPIClient struct {
	cc *grpc.ClientConn
}

func NewEchoAPIClient(cc *grpc.ClientConn) EchoAPIClient {
	return &echoAPIClient{cc}
}

func (c *echoAPIClient) Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error) {
	out := new(EchoResponse)
	err := c.cc.Invoke(ctx, "/service.EchoAPI/Echo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EchoAPIServer is the server API for EchoAPI service.
type EchoAPIServer interface {
	Echo(context.Context, *EchoRequest) (*EchoResponse, error)
}

func RegisterEchoAPIServer(s *grpc.Server, srv EchoAPIServer) {
	s.RegisterService(&_EchoAPI_serviceDesc, srv)
}

func _EchoAPI_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EchoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoAPIServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.EchoAPI/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoAPIServer).Echo(ctx, req.(*EchoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EchoAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.EchoAPI",
	HandlerType: (*EchoAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _EchoAPI_Echo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service/echo.proto",
}

func init() { proto.RegisterFile("service/echo.proto", fileDescriptor_echo_ff2fe5810199bc2f) }

var fileDescriptor_echo_ff2fe5810199bc2f = []byte{
	// 139 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x4f, 0x4d, 0xce, 0xc8, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62,
	0x87, 0x8a, 0x29, 0xa9, 0x73, 0x71, 0xbb, 0x26, 0x67, 0xe4, 0x07, 0xa5, 0x16, 0x96, 0xa6, 0x16,
	0x97, 0x08, 0x49, 0x70, 0xb1, 0xe7, 0xa6, 0x16, 0x17, 0x27, 0xa6, 0xa7, 0x4a, 0x30, 0x2a, 0x30,
	0x6a, 0x70, 0x06, 0xc1, 0xb8, 0x4a, 0x1a, 0x5c, 0x3c, 0x10, 0x85, 0xc5, 0x05, 0xf9, 0x79, 0xc5,
	0xa9, 0xb8, 0x55, 0x1a, 0x39, 0x70, 0xb1, 0x83, 0x54, 0x3a, 0x06, 0x78, 0x0a, 0x99, 0x72, 0xb1,
	0x80, 0x98, 0x42, 0x22, 0x7a, 0x50, 0xfb, 0xf4, 0x90, 0x2c, 0x93, 0x12, 0x45, 0x13, 0x85, 0x98,
	0xac, 0xc4, 0xe0, 0xc4, 0x19, 0x05, 0x73, 0x5f, 0x12, 0x1b, 0xd8, 0xbd, 0xc6, 0x80, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x08, 0x0b, 0x5b, 0xa4, 0xc5, 0x00, 0x00, 0x00,
}
