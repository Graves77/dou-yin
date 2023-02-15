package svc

import (
	"awesomeProject/dou-yin/apps/user/cmd/api/internal/config"
	"awesomeProject/dou-yin/apps/user/cmd/rpc/client"
	"awesomeProject/dou-yin/apps/user/cmd/rpc/types/user"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	Redis   *redis.Redis
	UserRpc user.RpcClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: client.NewRpc(zrpc.MustNewClient(c.UserRpc)),
		Redis:   redis.New(c.Redis.Host),
	}
}
