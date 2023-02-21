package logic

import (
	"awesomeProject/dou-yin/apps/commaction/cmd/rpc/mysqlmanage/internal/svc"
	"awesomeProject/dou-yin/apps/commaction/cmd/rpc/mysqlmanage/types/mysqlmanageserver"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVideoListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetVideoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVideoListLogic {
	return &GetVideoListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 发布列表视频信息
func (l *GetVideoListLogic) GetVideoList(in *mysqlmanageserver.GetVideoListRequest) (*mysqlmanageserver.GetVideoListResponse, error) {
	// todo: add your logic here and delete this line

	return &mysqlmanageserver.GetVideoListResponse{}, nil
}
