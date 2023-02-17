package svc

import (
	"awesomeProject/dou-yin/video/cmd/api/internal/config"
	"awesomeProject/dou-yin/video/cmd/api/internal/middleware"
	"awesomeProject/dou-yin/video/cmd/rpc/publishvideo"
	"awesomeProject/dou-yin/video/cmd/rpc/types/video"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	UploadFile rest.Middleware
	Redis      *redis.Redis
	VideoRpc   video.PublishVideoClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		UploadFile: middleware.NewUploadFileMiddleware().Handle,
		Redis:      redis.New(c.Redis.Host),
		VideoRpc:   publishvideo.NewPublishVideo(zrpc.MustNewClient(c.VideoRpc)),
	}
}
