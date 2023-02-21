package svc

import (
	"awesomeProject/dou-yin/apps/user/cmd/rpc/types/user"
	"awesomeProject/dou-yin/apps/video/cmd/api/internal/config"
	"awesomeProject/dou-yin/apps/video/cmd/api/internal/middleware"
	"awesomeProject/dou-yin/apps/video/cmd/rpc/publishvideo"
	"awesomeProject/dou-yin/apps/video/cmd/rpc/types/video"
	"github.com/streadway/amqp"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type RabbitMQContext struct {
	Conn        *amqp.Connection
	Channel     *amqp.Channel
	ContentType string
}

type ServiceContext struct {
	Config     config.Config
	UploadFile rest.Middleware
	Redis      *redis.Redis
	VideoRpc   video.PublishVideoClient
	UserRpc    user.RpcClient
	RabbitMQ   *RabbitMQContext
}

func NewServiceContext(c config.Config) *ServiceContext {
	coon, err := amqp.Dial(c.RabbitMQ.GetUrl())
	if err != nil {
		//panic(err)
	}
	channel, err := coon.Channel()
	if err != nil {
		//panic(err)
	}

	rabbitMQ := &RabbitMQContext{
		Conn:        coon,
		Channel:     channel,
		ContentType: "text/plain",
	}

	return &ServiceContext{
		Config:     c,
		UploadFile: middleware.NewUploadFileMiddleware().Handle,
		Redis:      redis.New(c.Redis.Host),
		VideoRpc:   publishvideo.NewPublishVideo(zrpc.MustNewClient(c.VideoRpc)),
		RabbitMQ:   rabbitMQ,
	}
}
