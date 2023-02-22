package svc

import (
	"awesomeProject/dou-yin/apps/api/internal/config"
	"awesomeProject/dou-yin/apps/api/internal/middleware"
	"awesomeProject/dou-yin/apps/commaction/rpc/mongodbmanage/mongodbmanageserverclient"
	"awesomeProject/dou-yin/apps/commaction/rpc/mysqlmanage/mysqlmanageserverclient"
	"awesomeProject/dou-yin/apps/user/cmd/rpc/types/user"
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
	Config           config.Config
	MySQLManageRpc   mysqlmanageserverclient.MySQLManageServer
	MongoDBMangerRpc mongodbmanageserverclient.MongodbManageServer
	UploadFile       rest.Middleware
	Redis            *redis.Redis
	VideoRpc         video.PublishVideoClient
	UserRpc          user.RpcClient
	RabbitMQ         *RabbitMQContext
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
		Config:           c,
		MySQLManageRpc:   mysqlmanageserverclient.NewMySQLManageServer(zrpc.MustNewClient(c.MySQLManageRpc)),
		MongoDBMangerRpc: mongodbmanageserverclient.NewMongodbManageServer(zrpc.MustNewClient(c.MongoDBManageRpc)),
		UploadFile:       middleware.NewUploadFileMiddleware().Handle,
		Redis:            redis.New(c.Redis.Host),
		VideoRpc:         publishvideo.NewPublishVideo(zrpc.MustNewClient(c.VideoRpc)),
		RabbitMQ:         rabbitMQ,
	}
}
