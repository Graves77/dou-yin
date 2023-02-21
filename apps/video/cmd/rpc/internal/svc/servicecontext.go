package svc

import (
	"awesomeProject/dou-yin/apps/video/cmd/rpc/internal/config"
	"awesomeProject/dou-yin/apps/video/model"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config     config.Config
	VideoModel model.VideosModel
	Redis      *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:     c,
		VideoModel: model.NewVideosModel(sqlConn, c.CacheRedis),
		Redis:      redis.New(c.Redis.Host),
	}
}
