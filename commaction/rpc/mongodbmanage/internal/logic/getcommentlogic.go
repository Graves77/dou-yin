package logic

import (
	"context"

	"awesomeProject/dou-yin/commaction/rpc/mongodbmanage/internal/svc"
	"awesomeProject/dou-yin/commaction/rpc/mongodbmanage/types/mongodbmanageserver"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentLogic {
	return &GetCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCommentLogic) GetComment(in *mongodbmanageserver.CommentListRequest) (*mongodbmanageserver.CommentListResponse, error) {
	// todo: add your logic here and delete this line

	return &mongodbmanageserver.CommentListResponse{}, nil
}
