package logic

import (
	"context"

	"awesomeProject/dou-yin/commaction/rpc/mongodbmanage/internal/svc"
	"awesomeProject/dou-yin/commaction/rpc/mongodbmanage/types/mongodbmanageserver"

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
