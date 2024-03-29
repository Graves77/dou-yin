package logic

import (
	"awesomeProject/dou-yin/apps/commaction/rpc/mysqlmanage/internal/svc"
	"awesomeProject/dou-yin/apps/commaction/rpc/mysqlmanage/types/mysqlmanageserver"
	"context"

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
func (l *CheckUserInfLogic) CheckUserInf(in *mysqlmanageserver.CheckUserInfRequest) (*mysqlmanageserver.CheckIsFollowResponse, error) {
	var num int64
	err := svc.DB.Table("follow_and_follower_list").Where("user_id = ? and follower_id = ?", in.UserId, in.FollowerId).Count(&num).Error
	if err != nil {
		logx.Errorf("[pkg]logic [func]CheckIsFollow [msg]gorm follow_and_follower_list.Take %v", err)
		return &mysqlmanageserver.CheckIsFollowResponse{Ok: false}, err
	}
	return &mysqlmanageserver.CheckIsFollowResponse{Ok: num > 0}, nil
}
