package svc

import (
	"awesomeProject/dou-yin/like/cmd/mq/internal/config"
	"awesomeProject/dou-yin/like/cmd/rpc/likeclient"
	"awesomeProject/dou-yin/like/cmd/rpc/types/like"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	LikeRpc like.LikeClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		LikeRpc: likeclient.NewLike(zrpc.MustNewClient(c.LikeRpc)),
	}
}
