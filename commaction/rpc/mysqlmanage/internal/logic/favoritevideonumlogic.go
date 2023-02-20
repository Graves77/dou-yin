package logic

import (
	"context"

	"awesomeProject/dou-yin/commaction/rpc/mysqlmanage/internal/svc"
	"awesomeProject/dou-yin/commaction/rpc/mysqlmanage/types/mysqlmanageserver"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteVideoNumLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFavoriteVideoNumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteVideoNumLogic {
	return &FavoriteVideoNumLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 喜欢视频数量
func (l *FavoriteVideoNumLogic) FavoriteVideoNum(in *mysqlmanageserver.FavoriteVideoNumRequest) (*mysqlmanageserver.FavoriteVideoNumResponse, error) {
	// todo: add your logic here and delete this line

	return &mysqlmanageserver.FavoriteVideoNumResponse{}, nil
}
