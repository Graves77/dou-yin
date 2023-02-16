package svc

import (
	"awesomeProject/dou-yin/video/cmd/rpc/internal/config"
	"awesomeProject/dou-yin/video/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config     config.Config
	VideoModel model.VideosModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:     c,
		VideoModel: model.NewVideosModel(sqlConn, c.CacheRedis),
	}
}
