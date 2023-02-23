package logic

import (
	"awesomeProject/dou-yin/apps/commaction/rpc/mysqlmanage/internal/svc"
	"awesomeProject/dou-yin/apps/commaction/rpc/mysqlmanage/types/mysqlmanageserver"
	"context"
	"gorm.io/gorm"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddVideoFavoriteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}
type Favorite_list struct {
	Favorite_video_id int64
	Favorite_user_id  int
	Record_time       time.Time
}

func NewAddVideoFavoriteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddVideoFavoriteLogic {
	return &AddVideoFavoriteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 点赞
func (l *AddVideoFavoriteLogic) AddVideoFavorite(in *mysqlmanageserver.AddVideoFavoriteRequest) (*mysqlmanageserver.AddVideoFavoriteResponse, error) {
	var n int64
	err := svc.DB.Table("favorite_list").Where("favorite_video_id = ? AND favorite_user_id = ?", in.VideoID, in.UserID).Count(&n).Error
	if err != nil {
		logx.Errorf("[pkg]logic [func]AddVideoFavorite [msg]gorm favorite_list.Count [err]%v", err)
		return &mysqlmanageserver.AddVideoFavoriteResponse{Ok: 0}, err
	}
	if int(n) > 0 {
		logx.Infof("[pkg]logic [func]AddVideoFavorite [msg]gorm favorite already")
		return &mysqlmanageserver.AddVideoFavoriteResponse{Ok: -2}, nil
	}

	err = svc.DB.Table("video_info").Where("video_id = ?", in.VideoID).Count(&n).Error
	if err != nil {
		logx.Errorf("[pkg]logic [func]AddVideoFavorite [msg]gorm favorite_list.Count [err]%v", err)
		return &mysqlmanageserver.AddVideoFavoriteResponse{Ok: 0}, nil
	}
	if int(n) == 0 {
		logx.Infof("[pkg]logic [func]AddVideoFavorite [msg]Video does no exist")
		return &mysqlmanageserver.AddVideoFavoriteResponse{Ok: -1}, nil
	}

	err = svc.DB.Table("video_info").Where("video_id = ?", in.VideoID).Update("favorite_count", gorm.Expr("favorite_count + 1")).Error
	if err != nil {
		logx.Errorf("[pkg]logic [func]AddVideoFavorite [msg]gorm favorite_count.UPDATE [err]%v", err)
		return &mysqlmanageserver.AddVideoFavoriteResponse{Ok: 0}, nil
	}

	fl := Favorite_list{
		Favorite_video_id: in.VideoID,
		Favorite_user_id:  int(in.UserID),
		Record_time:       time.Now(),
	}

	err = svc.DB.Table("favorite_list").Omit("record_id").Create(&fl).Error
	if err != nil {
		logx.Errorf("[pkg]logic [func]AddVideoFavorite [msg]gorm favorite_list.Create [err]%v", err)
		return &mysqlmanageserver.AddVideoFavoriteResponse{Ok: 0}, nil
	}
	return &mysqlmanageserver.AddVideoFavoriteResponse{Ok: 1}, err
}
