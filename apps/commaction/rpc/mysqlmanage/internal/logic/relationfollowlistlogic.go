package logic

import (
	"awesomeProject/dou-yin/apps/commaction/rpc/mysqlmanage/internal/svc"
	"awesomeProject/dou-yin/apps/commaction/rpc/mysqlmanage/types/mysqlmanageserver"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type RelationFollowListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRelationFollowListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RelationFollowListLogic {
	return &RelationFollowListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 关注列表
func (l *RelationFollowListLogic) RelationFollowList(in *mysqlmanageserver.RelationFollowListRequest) (*mysqlmanageserver.RelationFollowListResponse, error) {
	// todo: add your logic here and delete this line

	return &mysqlmanageserver.RelationFollowListResponse{}, nil
}
