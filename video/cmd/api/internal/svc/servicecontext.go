package svc

import (
	"awesomeProject/dou-yin/video/cmd/api/internal/config"
	"awesomeProject/dou-yin/video/cmd/api/internal/middleware"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config     config.Config
	UploadFile rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		UploadFile: middleware.NewUploadFileMiddleware().Handle,
	}
}
