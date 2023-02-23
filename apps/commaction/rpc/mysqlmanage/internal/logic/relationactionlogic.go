package logic

import (
	"awesomeProject/dou-yin/apps/commaction/rpc/mysqlmanage/internal/svc"
	"awesomeProject/dou-yin/apps/commaction/rpc/mysqlmanage/types/mysqlmanageserver"
	"context"
	"errors"
	"gorm.io/gorm"
	"time"

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

type follow_and_follower_list struct {
	UserID     int64     `gorm:"cloumn:user_id;"`
	FollowerId int64     `gorm:"cloumn:follower_id;"`
	RecordTime time.Time `gorm:"cloumn:record_time;"`
}

// 关注、取消关注
func (l *RelationActionLogic) RelationAction(in *mysqlmanageserver.RelationActionRequest) (*mysqlmanageserver.RelationActionResponse, error) {
	db := svc.DB
	if in.ActionType == 1 { //关注
		//#关注账号 user_id：被关注的账号  follower_id：哪个账号要关注
		insertData := follow_and_follower_list{UserID: in.ToUserID, FollowerId: in.UserID, RecordTime: time.Now()}
		if err := db.Table("follow_and_follower_list").Create(&insertData).Error; err != nil {
			return &mysqlmanageserver.RelationActionResponse{Ok: false}, err
		}
		if err := db.Table("user_info").Where("user_id = ?", in.ToUserID).Update("follower_count", gorm.Expr("follower_count + ?", 1)).Error; err != nil {
			return &mysqlmanageserver.RelationActionResponse{Ok: false}, nil
		}
		if err := db.Table("user_info").Where("user_id = ?", in.UserID).Update("follow_count", gorm.Expr("follow_count + ?", 1)).Error; err != nil {
			return &mysqlmanageserver.RelationActionResponse{Ok: false}, nil
		}
		return &mysqlmanageserver.RelationActionResponse{Ok: true}, nil
	} else if in.ActionType == 2 { //取消关注
		//取消关注.  user_id：要被取消关注的账号   follower_id：哪个账号要取消关注

		if err := db.Table("follow_and_follower_list").Where("user_id = ? and follower_id = ?", in.ToUserID, in.UserID).Delete(nil).Error; err != nil {
			return &mysqlmanageserver.RelationActionResponse{Ok: false}, err
		}
		if err := db.Table("user_info").Where("user_id = ?", in.ToUserID).Update("follower_count", gorm.Expr("follower_count - ?", 1)).Error; err != nil {
			return &mysqlmanageserver.RelationActionResponse{Ok: false}, nil
		}
		if err := db.Table("user_info").Where("user_id = ?", in.UserID).Update("follow_count", gorm.Expr("follow_count - ?", 1)).Error; err != nil {
			return &mysqlmanageserver.RelationActionResponse{Ok: false}, nil
		}
	} else {
		return &mysqlmanageserver.RelationActionResponse{Ok: false}, errors.New("无效动作")
	}

	return &mysqlmanageserver.RelationActionResponse{Ok: true}, nil
}
