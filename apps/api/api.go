package main

import (
	"awesomeProject/dou-yin/apps/api/internal/config"
	"awesomeProject/dou-yin/apps/api/internal/handler"
	"awesomeProject/dou-yin/apps/api/internal/svc"
	"awesomeProject/dou-yin/common/result"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf, rest.WithUnauthorizedCallback(func(w http.ResponseWriter, r *http.Request, err error) {
		httpx.WriteJson(w, http.StatusOK, result.Fail("JWT Auth Failed"))
	}))
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
