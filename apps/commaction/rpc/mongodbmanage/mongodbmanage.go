package main

import (
	"awesomeProject/dou-yin/apps/commaction/rpc/mongodbmanage/internal/config"
	"awesomeProject/dou-yin/apps/commaction/rpc/mongodbmanage/internal/server"
	"awesomeProject/dou-yin/apps/commaction/rpc/mongodbmanage/internal/svc"
	"awesomeProject/dou-yin/apps/commaction/rpc/mongodbmanage/types/mongodbmanageserver"
	"flag"
	"fmt"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/mongodbmanage.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		mongodbmanageserver.RegisterMongodbManageServerServer(grpcServer, server.NewMongodbManageServerServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
