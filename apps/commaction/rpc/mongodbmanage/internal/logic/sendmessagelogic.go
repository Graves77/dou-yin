package logic

import (
	"awesomeProject/dou-yin/apps/commaction/rpc/mongodbmanage/internal/svc"
	"awesomeProject/dou-yin/apps/commaction/rpc/mongodbmanage/types/mongodbmanageserver"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type Message struct {
	Id         int64  `json:"id"                    form:"id"                    bson:"_id"`
	ToUserId   int64  `json:"to_user_id"            form:"to_user_id"            bson:"to_user_id"`
	FromUserId int64  `json:"from_user_id"          form:"from_user_id"          bson:"from_user_id"`
	Content    string `json:"content"               form:"content"               bson:"content"`
	CreateTime string `json:"create_time"           form:"create_time"           bson:"create_time"`
}

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
