package logic

import (
	"awesomeProject/dou-yin/apps/commaction/rpc/mysqlmanage/internal/svc"
	"awesomeProject/dou-yin/apps/commaction/rpc/mysqlmanage/types/mysqlmanageserver"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFavoriteVideoListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFavoriteVideoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFavoriteVideoListLogic {
	return &GetFavoriteVideoListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取喜欢列表
func (l *GetFavoriteVideoListLogic) GetFavoriteVideoList(in *mysqlmanageserver.GetFavoriteVideoListRequest) (*mysqlmanageserver.GetFavoriteVideoListResponse, error) {
	// todo: add your logic here and delete this line

	return &mysqlmanageserver.GetFavoriteVideoListResponse{}, nil
}
