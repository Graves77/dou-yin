package logic

import (
	"context"

	"awesomeProject/dou-yin/commaction/rpc/mysqlmanage/internal/svc"
	"awesomeProject/dou-yin/commaction/rpc/mysqlmanage/types/mysqlmanageserver"

	"github.com/zeromicro/go-zero/core/logx"
)

type VideoNumLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewVideoNumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VideoNumLogic {
	return &VideoNumLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 发布列表视频数量
func (l *VideoNumLogic) VideoNum(in *mysqlmanageserver.VideoNumRequest) (*mysqlmanageserver.VideoNumResponse, error) {
	// todo: add your logic here and delete this line

	return &mysqlmanageserver.VideoNumResponse{}, nil
}
