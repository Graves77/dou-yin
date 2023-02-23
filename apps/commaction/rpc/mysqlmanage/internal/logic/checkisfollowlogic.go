package logic

import (
	"awesomeProject/dou-yin/apps/commaction/rpc/mysqlmanage/internal/svc"
	"awesomeProject/dou-yin/apps/commaction/rpc/mysqlmanage/types/mysqlmanageserver"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckIsFollowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckIsFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckIsFollowLogic {
	return &CheckIsFollowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 是否关注
func (l *CheckIsFollowLogic) CheckIsFollow(in *mysqlmanageserver.CheckIsFollowRequest) (*mysqlmanageserver.CheckIsFollowResponse, error) {

	var num int64
	err := svc.DB.Table("follow_and_follower_list").Where("user_id = ? and follower_id = ?", in.UserId, in.FollowerId).Count(&num).Error
	if err != nil {
		logx.Errorf("[pkg]logic [func]CheckIsFollow [msg]gorm follow_and_follower_list.Take %v", err)
		return &mysqlmanageserver.CheckIsFollowResponse{Ok: false}, err
	}
	return &mysqlmanageserver.CheckIsFollowResponse{Ok: num > 0}, nil
}
