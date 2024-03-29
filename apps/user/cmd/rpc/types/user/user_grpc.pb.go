// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: user.proto

package user

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RpcClient is the client API for Rpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RpcClient interface {
	Login(ctx context.Context, in *DouyinUserLoginRequest, opts ...grpc.CallOption) (*DouyinUserLoginResponse, error)
	Register(ctx context.Context, in *DouyinUserRegisterRequest, opts ...grpc.CallOption) (*DouyinUserRegisterResponse, error)
	UserInfo(ctx context.Context, in *DouyinUserRequest, opts ...grpc.CallOption) (*DouyinUserResponse, error)
}

type rpcClient struct {
	cc grpc.ClientConnInterface
}

func NewRpcClient(cc grpc.ClientConnInterface) RpcClient {
	return &rpcClient{cc}
}

func (c *rpcClient) Login(ctx context.Context, in *DouyinUserLoginRequest, opts ...grpc.CallOption) (*DouyinUserLoginResponse, error) {
	out := new(DouyinUserLoginResponse)
	err := c.cc.Invoke(ctx, "/user.Rpc/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcClient) Register(ctx context.Context, in *DouyinUserRegisterRequest, opts ...grpc.CallOption) (*DouyinUserRegisterResponse, error) {
	out := new(DouyinUserRegisterResponse)
	err := c.cc.Invoke(ctx, "/user.Rpc/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcClient) UserInfo(ctx context.Context, in *DouyinUserRequest, opts ...grpc.CallOption) (*DouyinUserResponse, error) {
	out := new(DouyinUserResponse)
	err := c.cc.Invoke(ctx, "/user.Rpc/UserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RpcServer is the server API for Rpc service.
// All implementations must embed UnimplementedRpcServer
// for forward compatibility
type RpcServer interface {
	Login(context.Context, *DouyinUserLoginRequest) (*DouyinUserLoginResponse, error)
	Register(context.Context, *DouyinUserRegisterRequest) (*DouyinUserRegisterResponse, error)
	UserInfo(context.Context, *DouyinUserRequest) (*DouyinUserResponse, error)
	mustEmbedUnimplementedRpcServer()
}

// UnimplementedRpcServer must be embedded to have forward compatible implementations.
type UnimplementedRpcServer struct {
}

func (UnimplementedRpcServer) Login(context.Context, *DouyinUserLoginRequest) (*DouyinUserLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedRpcServer) Register(context.Context, *DouyinUserRegisterRequest) (*DouyinUserRegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedRpcServer) UserInfo(context.Context, *DouyinUserRequest) (*DouyinUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserInfo not implemented")
}
func (UnimplementedRpcServer) mustEmbedUnimplementedRpcServer() {}

// UnsafeRpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RpcServer will
// result in compilation errors.
type UnsafeRpcServer interface {
	mustEmbedUnimplementedRpcServer()
}

func RegisterRpcServer(s grpc.ServiceRegistrar, srv RpcServer) {
	s.RegisterService(&Rpc_ServiceDesc, srv)
}

func _Rpc_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DouyinUserLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.Rpc/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServer).Login(ctx, req.(*DouyinUserLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rpc_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DouyinUserRegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.Rpc/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServer).Register(ctx, req.(*DouyinUserRegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rpc_UserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DouyinUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServer).UserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.Rpc/UserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServer).UserInfo(ctx, req.(*DouyinUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Rpc_ServiceDesc is the grpc.ServiceDesc for Rpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Rpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.Rpc",
	HandlerType: (*RpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Rpc_Login_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _Rpc_Register_Handler,
		},
		{
			MethodName: "UserInfo",
			Handler:    _Rpc_UserInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
