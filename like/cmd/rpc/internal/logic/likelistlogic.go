package logic

import (
	"context"

	"awesomeProject/dou-yin/like/cmd/rpc/internal/svc"
	"awesomeProject/dou-yin/like/cmd/rpc/types/like"

	"github.com/zeromicro/go-zero/core/logx"
)

type LikeListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLikeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LikeListLogic {
	return &LikeListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LikeListLogic) LikeList(in *like.LikeListReq) (*like.LikeListResp, error) {
	// todo: add your logic here and delete this line

	return &like.LikeListResp{}, nil
}
