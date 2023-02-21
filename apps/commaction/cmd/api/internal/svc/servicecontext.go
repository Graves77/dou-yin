package svc

import (
	"awesomeProject/dou-yin/apps/commaction/cmd/api/internal/config"
	"awesomeProject/dou-yin/apps/commaction/cmd/rpc/mongodbmanage/mongodbmanageserverclient"
	"awesomeProject/dou-yin/apps/commaction/cmd/rpc/mysqlmanage/mysqlmanageserverclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config           config.Config
	MySQLManageRpc   mysqlmanageserverclient.MySQLManageServer
	MongoDBMangerRpc mongodbmanageserverclient.MongodbManageServer
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:           c,
		MySQLManageRpc:   mysqlmanageserverclient.NewMySQLManageServer(zrpc.MustNewClient(c.MySQLManageRpc)),
		MongoDBMangerRpc: mongodbmanageserverclient.NewMongodbManageServer(zrpc.MustNewClient(c.MongoDBManageRpc)),
	}
}
