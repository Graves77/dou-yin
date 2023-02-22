package logic

import (
	"awesomeProject/dou-yin/apps/commaction/rpc/mysqlmanage/internal/svc"
	"awesomeProject/dou-yin/apps/commaction/rpc/mysqlmanage/types/mysqlmanageserver"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type RelationFollowerListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRelationFollowerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RelationFollowerListLogic {
	return &RelationFollowerListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 粉丝列表
func (l *RelationFollowerListLogic) RelationFollowerList(in *mysqlmanageserver.RelationFollowerListRequest) (*mysqlmanageserver.RelationFollowerListResponse, error) {
	// todo: add your logic here and delete this line

	return &mysqlmanageserver.RelationFollowerListResponse{}, nil
}
