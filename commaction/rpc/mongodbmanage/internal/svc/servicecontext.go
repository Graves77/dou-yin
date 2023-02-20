package svc

import "awesomeProject/dou-yin/commaction/rpc/mongodbmanage/internal/config"

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
