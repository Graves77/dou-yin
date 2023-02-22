package logic

import (
	"awesomeProject/dou-yin/apps/commaction/rpc/mysqlmanage/internal/svc"
	"awesomeProject/dou-yin/apps/commaction/rpc/mysqlmanage/types/mysqlmanageserver"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type RelationActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRelationActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RelationActionLogic {
	return &RelationActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 关注、取消关注
func (l *RelationActionLogic) RelationAction(in *mysqlmanageserver.RelationActionRequest) (*mysqlmanageserver.RelationActionResponse, error) {
	// todo: add your logic here and delete this line

	return &mysqlmanageserver.RelationActionResponse{}, nil
}
