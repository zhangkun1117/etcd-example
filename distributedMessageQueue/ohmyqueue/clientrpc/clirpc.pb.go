// Code generated by protoc-gen-go.
// source: clirpc.proto
// DO NOT EDIT!

/*
Package clientrpc is a generated protocol buffer package.

It is generated from these files:
	clirpc.proto

It has these top-level messages:
	Req
	Msg
	Resp
	StatusCode
*/
package clientrpc

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

type Req struct {
	Topic  string `protobuf:"bytes,1,opt,name=topic" json:"topic,omitempty"`
	Offset int64  `protobuf:"varint,2,opt,name=offset" json:"offset,omitempty"`
}

func (m *Req) Reset()                    { *m = Req{} }
func (m *Req) String() string            { return proto.CompactTextString(m) }
func (*Req) ProtoMessage()               {}
func (*Req) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Req) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *Req) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type Msg struct {
	Topic string `protobuf:"bytes,1,opt,name=topic" json:"topic,omitempty"`
	Body  string `protobuf:"bytes,2,opt,name=body" json:"body,omitempty"`
}

func (m *Msg) Reset()                    { *m = Msg{} }
func (m *Msg) String() string            { return proto.CompactTextString(m) }
func (*Msg) ProtoMessage()               {}
func (*Msg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Msg) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *Msg) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

type Resp struct {
	Body   string `protobuf:"bytes,1,opt,name=body" json:"body,omitempty"`
	Offset int64  `protobuf:"varint,2,opt,name=offset" json:"offset,omitempty"`
}

func (m *Resp) Reset()                    { *m = Resp{} }
func (m *Resp) String() string            { return proto.CompactTextString(m) }
func (*Resp) ProtoMessage()               {}
func (*Resp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Resp) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *Resp) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type StatusCode struct {
	Code int32 `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
}

func (m *StatusCode) Reset()                    { *m = StatusCode{} }
func (m *StatusCode) String() string            { return proto.CompactTextString(m) }
func (*StatusCode) ProtoMessage()               {}
func (*StatusCode) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *StatusCode) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func init() {
	proto.RegisterType((*Req)(nil), "clientrpc.Req")
	proto.RegisterType((*Msg)(nil), "clientrpc.Msg")
	proto.RegisterType((*Resp)(nil), "clientrpc.Resp")
	proto.RegisterType((*StatusCode)(nil), "clientrpc.StatusCode")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Omq service

type OmqClient interface {
	PutMsg(ctx context.Context, in *Msg, opts ...grpc.CallOption) (*StatusCode, error)
	Poll(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Resp, error)
}

type omqClient struct {
	cc *grpc.ClientConn
}

func NewOmqClient(cc *grpc.ClientConn) OmqClient {
	return &omqClient{cc}
}

func (c *omqClient) PutMsg(ctx context.Context, in *Msg, opts ...grpc.CallOption) (*StatusCode, error) {
	out := new(StatusCode)
	err := grpc.Invoke(ctx, "/clientrpc.Omq/PutMsg", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *omqClient) Poll(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Resp, error) {
	out := new(Resp)
	err := grpc.Invoke(ctx, "/clientrpc.Omq/poll", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Omq service

type OmqServer interface {
	PutMsg(context.Context, *Msg) (*StatusCode, error)
	Poll(context.Context, *Req) (*Resp, error)
}

func RegisterOmqServer(s *grpc.Server, srv OmqServer) {
	s.RegisterService(&_Omq_serviceDesc, srv)
}

func _Omq_PutMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Msg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OmqServer).PutMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clientrpc.Omq/PutMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OmqServer).PutMsg(ctx, req.(*Msg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Omq_Poll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OmqServer).Poll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clientrpc.Omq/Poll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OmqServer).Poll(ctx, req.(*Req))
	}
	return interceptor(ctx, in, info, handler)
}

var _Omq_serviceDesc = grpc.ServiceDesc{
	ServiceName: "clientrpc.Omq",
	HandlerType: (*OmqServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PutMsg",
			Handler:    _Omq_PutMsg_Handler,
		},
		{
			MethodName: "poll",
			Handler:    _Omq_Poll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "clirpc.proto",
}

func init() { proto.RegisterFile("clirpc.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 208 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xc1, 0x4e, 0x85, 0x30,
	0x10, 0x45, 0xa9, 0x05, 0x12, 0x26, 0x46, 0x93, 0x46, 0x0d, 0x61, 0x45, 0xba, 0xc2, 0x0d, 0x46,
	0xf8, 0x04, 0xd7, 0x44, 0x53, 0xbf, 0x40, 0x4a, 0x21, 0x24, 0xd5, 0x29, 0xb4, 0x2c, 0xfc, 0x7b,
	0x43, 0x7d, 0x81, 0xc7, 0x82, 0xdd, 0x3d, 0xb9, 0x39, 0xed, 0xcc, 0xc0, 0xad, 0xd4, 0xe3, 0x6c,
	0x64, 0x69, 0x66, 0x74, 0xc8, 0x12, 0xa9, 0x47, 0xf5, 0xe3, 0x66, 0x23, 0x79, 0x0d, 0x54, 0xa8,
	0x89, 0x3d, 0x40, 0xe4, 0xd0, 0x8c, 0x32, 0x25, 0x39, 0x29, 0x12, 0xf1, 0x0f, 0xec, 0x09, 0x62,
	0xec, 0x7b, 0xab, 0x5c, 0x7a, 0x93, 0x93, 0x82, 0x8a, 0x0b, 0xf1, 0x17, 0xa0, 0x8d, 0x1d, 0x4e,
	0x24, 0x06, 0x61, 0x8b, 0xdd, 0xaf, 0x57, 0x12, 0xe1, 0x33, 0xaf, 0x20, 0x14, 0xca, 0x9a, 0xad,
	0x23, 0x7b, 0x77, 0xfa, 0x49, 0x0e, 0xf0, 0xe9, 0xbe, 0xdc, 0x62, 0xdf, 0xb0, 0x53, 0xab, 0x29,
	0xb1, 0x53, 0xde, 0x8c, 0x84, 0xcf, 0x95, 0x04, 0xfa, 0xfe, 0x3d, 0xb1, 0x57, 0x88, 0x3f, 0x16,
	0xb7, 0x0e, 0x74, 0x57, 0x6e, 0x8b, 0x95, 0x8d, 0x1d, 0xb2, 0xc7, 0x2b, 0xde, 0xdf, 0xe2, 0x01,
	0x7b, 0x86, 0xd0, 0xa0, 0xd6, 0x07, 0x41, 0xa8, 0x29, 0xbb, 0x3f, 0xb0, 0x35, 0x3c, 0x68, 0x63,
	0x7f, 0xb2, 0xfa, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x82, 0x30, 0x8e, 0xfa, 0x42, 0x01, 0x00, 0x00,
}
