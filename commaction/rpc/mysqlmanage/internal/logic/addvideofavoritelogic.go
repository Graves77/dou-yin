package logic

import (
	"context"

	"awesomeProject/dou-yin/commaction/rpc/mysqlmanage/internal/svc"
	"awesomeProject/dou-yin/commaction/rpc/mysqlmanage/types/mysqlmanageserver"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddVideoFavoriteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddVideoFavoriteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddVideoFavoriteLogic {
	return &AddVideoFavoriteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 点赞
func (l *AddVideoFavoriteLogic) AddVideoFavorite(in *mysqlmanageserver.AddVideoFavoriteRequest) (*mysqlmanageserver.AddVideoFavoriteResponse, error) {
	// todo: add your logic here and delete this line

	return &mysqlmanageserver.AddVideoFavoriteResponse{}, nil
}
