package svc

import (
	"awesomeProject/dou-yin/like/cmd/rpc/internal/config"
	"awesomeProject/dou-yin/like/model"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config        config.Config
	FavoriteModel model.FavoriteModel
	Redis         *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		FavoriteModel: model.NewFavoriteModel(sqlx.NewMysql(c.Mysql.DataSource), c.CacheRedis),
		Redis:         redis.New(c.Redis.Host),
	}
}
