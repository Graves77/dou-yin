package logic

import (
	"awesomeProject/dou-yin/apps/commaction/cmd/rpc/mongodbmanage/internal/svc"
	"awesomeProject/dou-yin/apps/commaction/cmd/rpc/mongodbmanage/types/mongodbmanageserver"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMessageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMessageLogic {
	return &GetMessageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMessageLogic) GetMessage(in *mongodbmanageserver.MessageChatRequest) (*mongodbmanageserver.MessageChatResponse, error) {
	// todo: add your logic here and delete this line

	return &mongodbmanageserver.MessageChatResponse{}, nil
}
