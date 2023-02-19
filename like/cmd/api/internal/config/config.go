package config

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type MqConf struct {
	Username string
	Password string
	Host     string
	Port     int
}

func (c *MqConf) GetUrl() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%d/", c.Username, c.Password,
		c.Host, c.Port)
}

type Config struct {
	rest.RestConf
	UserRpc  zrpc.RpcClientConf
	VideoRpc zrpc.RpcClientConf
	JwtAuth  struct {
		AccessSecret string
		AccessExpire int64
	}
	Redis redis.RedisConf

	RabbitMQ MqConf
}
