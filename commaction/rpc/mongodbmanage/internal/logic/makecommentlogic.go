package logic

import (
	"context"

	"awesomeProject/dou-yin/commaction/rpc/mongodbmanage/internal/svc"
	"awesomeProject/dou-yin/commaction/rpc/mongodbmanage/types/mongodbmanageserver"

	"github.com/zeromicro/go-zero/core/logx"
)

type MakeCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMakeCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MakeCommentLogic {
	return &MakeCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MakeCommentLogic) MakeComment(in *mongodbmanageserver.CommentActionRequest) (*mongodbmanageserver.CommentActionResponse, error) {
	// todo: add your logic here and delete this line

	return &mongodbmanageserver.CommentActionResponse{}, nil
}
