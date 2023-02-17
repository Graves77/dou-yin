package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	VideoRpc zrpc.RpcClientConf
	JwtAuth  struct {
		AccessSecret string
		AccessExpire int64
	}
	UserRpc zrpc.RpcClientConf
	Redis   redis.RedisConf
}
