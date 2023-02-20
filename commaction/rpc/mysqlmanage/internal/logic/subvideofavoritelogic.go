package logic

import (
	"context"

	"awesomeProject/dou-yin/commaction/rpc/mysqlmanage/internal/svc"
	"awesomeProject/dou-yin/commaction/rpc/mysqlmanage/types/mysqlmanageserver"

	"github.com/zeromicro/go-zero/core/logx"
)

type SubVideoFavoriteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSubVideoFavoriteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubVideoFavoriteLogic {
	return &SubVideoFavoriteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 取消点赞
func (l *SubVideoFavoriteLogic) SubVideoFavorite(in *mysqlmanageserver.SubVideoFavoriteRequest) (*mysqlmanageserver.SubVideoFavoriteResponse, error) {
	// todo: add your logic here and delete this line

	return &mysqlmanageserver.SubVideoFavoriteResponse{}, nil
}
