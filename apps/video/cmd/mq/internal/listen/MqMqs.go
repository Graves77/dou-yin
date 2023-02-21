package listen

import (
	"awesomeProject/dou-yin/apps/video/cmd/mq/internal/svc"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/service"
)

func AddMqService(group *service.ServiceGroup, ctx context.Context, svcCtx *svc.ServiceContext) {
	listener := MustNewListener(svcCtx, handler)
	group.Add(listener)
}

func handler(svcCtx *svc.ServiceContext, message string) error {

	fmt.Println(message)

	return nil
}
