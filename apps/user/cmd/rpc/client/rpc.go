// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package client

import (
	"context"

	"awesomeProject/dou-yin/apps/user/cmd/rpc/types/user"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	DouyinUserLoginRequest     = user.DouyinUserLoginRequest
	DouyinUserLoginResponse    = user.DouyinUserLoginResponse
	DouyinUserRegisterRequest  = user.DouyinUserRegisterRequest
	DouyinUserRegisterResponse = user.DouyinUserRegisterResponse
	DouyinUserRequest          = user.DouyinUserRequest
	DouyinUserResponse         = user.DouyinUserResponse

	Rpc interface {
		Login(ctx context.Context, in *DouyinUserLoginRequest, opts ...grpc.CallOption) (*DouyinUserLoginResponse, error)
		Register(ctx context.Context, in *DouyinUserRegisterRequest, opts ...grpc.CallOption) (*DouyinUserRegisterResponse, error)
		UserInfo(ctx context.Context, in *DouyinUserRequest, opts ...grpc.CallOption) (*DouyinUserResponse, error)
	}

	defaultRpc struct {
		cli zrpc.Client
	}
)

func NewRpc(cli zrpc.Client) Rpc {
	return &defaultRpc{
		cli: cli,
	}
}

func (m *defaultRpc) Login(ctx context.Context, in *DouyinUserLoginRequest, opts ...grpc.CallOption) (*DouyinUserLoginResponse, error) {
	client := user.NewRpcClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

func (m *defaultRpc) Register(ctx context.Context, in *DouyinUserRegisterRequest, opts ...grpc.CallOption) (*DouyinUserRegisterResponse, error) {
	client := user.NewRpcClient(m.cli.Conn())
	return client.Register(ctx, in, opts...)
}

func (m *defaultRpc) UserInfo(ctx context.Context, in *DouyinUserRequest, opts ...grpc.CallOption) (*DouyinUserResponse, error) {
	client := user.NewRpcClient(m.cli.Conn())
	return client.UserInfo(ctx, in, opts...)
}