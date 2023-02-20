package svc

import (
	"awesomeProject/dou-yin/commaction/api/internal/config"
	"awesomeProject/dou-yin/commaction/rpc/mongodbmanage/mongodbmanageserverclient"
	"awesomeProject/dou-yin/commaction/rpc/mysqlmanage/mysqlmanageserverclient"
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
