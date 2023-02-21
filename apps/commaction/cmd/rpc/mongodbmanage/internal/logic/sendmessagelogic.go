package logic

import (
	"awesomeProject/dou-yin/apps/commaction/cmd/rpc/mongodbmanage/internal/svc"
	"awesomeProject/dou-yin/apps/commaction/cmd/rpc/mongodbmanage/types/mongodbmanageserver"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendMessageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendMessageLogic {
	return &SendMessageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendMessageLogic) SendMessage(in *mongodbmanageserver.MessageActionRequest) (*mongodbmanageserver.MessageActionResponse, error) {
	// todo: add your logic here and delete this line

	return &mongodbmanageserver.MessageActionResponse{}, nil
}
