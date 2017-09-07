// Code generated by protoc-gen-go. DO NOT EDIT.
// source: authlet.proto

/*
Package authlet is a generated protocol buffer package.

It is generated from these files:
	authlet.proto

It has these top-level messages:
	AuthRequest
	AuthReply
*/
package authlet

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

type AuthRequest struct {
	Clientid string `protobuf:"bytes,1,opt,name=clientid" json:"clientid,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password" json:"password,omitempty"`
	Topic    string `protobuf:"bytes,4,opt,name=topic" json:"topic,omitempty"`
	Hint     string `protobuf:"bytes,5,opt,name=hint" json:"hint,omitempty"`
	Access   string `protobuf:"bytes,6,opt,name=access" json:"access,omitempty"`
}

func (m *AuthRequest) Reset()                    { *m = AuthRequest{} }
func (m *AuthRequest) String() string            { return proto.CompactTextString(m) }
func (*AuthRequest) ProtoMessage()               {}
func (*AuthRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AuthRequest) GetClientid() string {
	if m != nil {
		return m.Clientid
	}
	return ""
}

func (m *AuthRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *AuthRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *AuthRequest) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *AuthRequest) GetHint() string {
	if m != nil {
		return m.Hint
	}
	return ""
}

func (m *AuthRequest) GetAccess() string {
	if m != nil {
		return m.Access
	}
	return ""
}

type AuthReply struct {
	Result  bool   `protobuf:"varint,1,opt,name=result" json:"result,omitempty"`
	Key     string `protobuf:"bytes,2,opt,name=key" json:"key,omitempty"`
	Version string `protobuf:"bytes,3,opt,name=version" json:"version,omitempty"`
}

func (m *AuthReply) Reset()                    { *m = AuthReply{} }
func (m *AuthReply) String() string            { return proto.CompactTextString(m) }
func (*AuthReply) ProtoMessage()               {}
func (*AuthReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AuthReply) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

func (m *AuthReply) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *AuthReply) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func init() {
	proto.RegisterType((*AuthRequest)(nil), "authlet.AuthRequest")
	proto.RegisterType((*AuthReply)(nil), "authlet.AuthReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Authlet service

type AuthletClient interface {
	// Get version of Authlet service
	GetVersion(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthReply, error)
	// Check acl based on client id, user name and topic
	CheckAcl(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthReply, error)
	// Check user name and password
	CheckUserNameAndPassword(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthReply, error)
	// Get PSK key
	GetPskKey(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthReply, error)
}

type authletClient struct {
	cc *grpc.ClientConn
}

func NewAuthletClient(cc *grpc.ClientConn) AuthletClient {
	return &authletClient{cc}
}

func (c *authletClient) GetVersion(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthReply, error) {
	out := new(AuthReply)
	err := grpc.Invoke(ctx, "/authlet.Authlet/GetVersion", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authletClient) CheckAcl(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthReply, error) {
	out := new(AuthReply)
	err := grpc.Invoke(ctx, "/authlet.Authlet/CheckAcl", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authletClient) CheckUserNameAndPassword(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthReply, error) {
	out := new(AuthReply)
	err := grpc.Invoke(ctx, "/authlet.Authlet/CheckUserNameAndPassword", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authletClient) GetPskKey(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthReply, error) {
	out := new(AuthReply)
	err := grpc.Invoke(ctx, "/authlet.Authlet/GetPskKey", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Authlet service

type AuthletServer interface {
	// Get version of Authlet service
	GetVersion(context.Context, *AuthRequest) (*AuthReply, error)
	// Check acl based on client id, user name and topic
	CheckAcl(context.Context, *AuthRequest) (*AuthReply, error)
	// Check user name and password
	CheckUserNameAndPassword(context.Context, *AuthRequest) (*AuthReply, error)
	// Get PSK key
	GetPskKey(context.Context, *AuthRequest) (*AuthReply, error)
}

func RegisterAuthletServer(s *grpc.Server, srv AuthletServer) {
	s.RegisterService(&_Authlet_serviceDesc, srv)
}

func _Authlet_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthletServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authlet.Authlet/GetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthletServer).GetVersion(ctx, req.(*AuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authlet_CheckAcl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthletServer).CheckAcl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authlet.Authlet/CheckAcl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthletServer).CheckAcl(ctx, req.(*AuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authlet_CheckUserNameAndPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthletServer).CheckUserNameAndPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authlet.Authlet/CheckUserNameAndPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthletServer).CheckUserNameAndPassword(ctx, req.(*AuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authlet_GetPskKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthletServer).GetPskKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authlet.Authlet/GetPskKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthletServer).GetPskKey(ctx, req.(*AuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Authlet_serviceDesc = grpc.ServiceDesc{
	ServiceName: "authlet.Authlet",
	HandlerType: (*AuthletServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetVersion",
			Handler:    _Authlet_GetVersion_Handler,
		},
		{
			MethodName: "CheckAcl",
			Handler:    _Authlet_CheckAcl_Handler,
		},
		{
			MethodName: "CheckUserNameAndPassword",
			Handler:    _Authlet_CheckUserNameAndPassword_Handler,
		},
		{
			MethodName: "GetPskKey",
			Handler:    _Authlet_GetPskKey_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "authlet.proto",
}

func init() { proto.RegisterFile("authlet.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 281 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x92, 0x3f, 0x4b, 0xf4, 0x40,
	0x10, 0xc6, 0xdf, 0xbc, 0x77, 0x97, 0x3f, 0x23, 0x82, 0x0c, 0x87, 0x2c, 0x57, 0x49, 0x2a, 0xab,
	0x2b, 0x14, 0xd4, 0x36, 0x08, 0x5e, 0x21, 0xe8, 0x11, 0xd0, 0x3e, 0x6e, 0x06, 0x12, 0xb2, 0x97,
	0xc4, 0xdd, 0x89, 0x92, 0xcf, 0xe3, 0xd7, 0xb4, 0x90, 0xec, 0x26, 0x87, 0xd8, 0xa5, 0x9b, 0xdf,
	0xfc, 0x32, 0xe1, 0xe1, 0x61, 0xe1, 0x34, 0xeb, 0xb8, 0x50, 0xc4, 0xdb, 0x56, 0x37, 0xdc, 0x60,
	0x30, 0x62, 0xfc, 0xe5, 0xc1, 0x49, 0xd2, 0x71, 0x91, 0xd2, 0x7b, 0x47, 0x86, 0x71, 0x03, 0xa1,
	0x54, 0x25, 0xd5, 0x5c, 0xe6, 0xc2, 0xbb, 0xf0, 0x2e, 0xa3, 0xf4, 0xc8, 0x83, 0xeb, 0x0c, 0xe9,
	0x3a, 0x3b, 0x90, 0xf8, 0xef, 0xdc, 0xc4, 0x83, 0x6b, 0x33, 0x63, 0x3e, 0x1b, 0x9d, 0x8b, 0x85,
	0x73, 0x13, 0xe3, 0x1a, 0x56, 0xdc, 0xb4, 0xa5, 0x14, 0x4b, 0x2b, 0x1c, 0x20, 0xc2, 0xb2, 0x28,
	0x6b, 0x16, 0x2b, 0xbb, 0xb4, 0x33, 0x9e, 0x83, 0x9f, 0x49, 0x49, 0xc6, 0x08, 0xdf, 0x6e, 0x47,
	0x8a, 0x9f, 0x21, 0x72, 0x21, 0x5b, 0xd5, 0x0f, 0x1f, 0x69, 0x32, 0x9d, 0x62, 0x1b, 0x30, 0x4c,
	0x47, 0xc2, 0x33, 0x58, 0x54, 0xd4, 0x8f, 0xc9, 0x86, 0x11, 0x05, 0x04, 0x1f, 0xa4, 0x4d, 0xd9,
	0xd4, 0x63, 0xa6, 0x09, 0xaf, 0xbe, 0x3d, 0x08, 0x12, 0x57, 0x01, 0xde, 0x01, 0xec, 0x88, 0x5f,
	0x9d, 0xc1, 0xf5, 0x76, 0x6a, 0xea, 0x57, 0x2d, 0x1b, 0xfc, 0xb3, 0x6d, 0x55, 0x1f, 0xff, 0xc3,
	0x1b, 0x08, 0xef, 0x0b, 0x92, 0x55, 0x22, 0xd5, 0xac, 0xbb, 0x07, 0x10, 0xf6, 0xee, 0xc5, 0x90,
	0x7e, 0xca, 0x0e, 0x94, 0xd4, 0xf9, 0xfe, 0x58, 0xd6, 0x8c, 0xff, 0xdc, 0x42, 0xb4, 0x23, 0xde,
	0x9b, 0xea, 0x91, 0xfa, 0x39, 0x87, 0x6f, 0xbe, 0x7d, 0x05, 0xd7, 0x3f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x53, 0xe5, 0xd6, 0x42, 0x16, 0x02, 0x00, 0x00,
}