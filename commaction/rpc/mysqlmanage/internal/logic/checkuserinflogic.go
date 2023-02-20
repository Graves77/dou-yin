package logic

import (
	"context"

	"awesomeProject/dou-yin/commaction/rpc/mysqlmanage/internal/svc"
	"awesomeProject/dou-yin/commaction/rpc/mysqlmanage/types/mysqlmanageserver"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckUserInfLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckUserInfLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckUserInfLogic {
	return &CheckUserInfLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获得用户信息
func (l *CheckUserInfLogic) CheckUserInf(in *mysqlmanageserver.CheckUserInfRequest) (*mysqlmanageserver.CheckUserInfResponse, error) {
	// todo: add your logic here and delete this line

	return &mysqlmanageserver.CheckUserInfResponse{}, nil
}
