package logic

import (
	"awesomeProject/dou-yin/apps/commaction/rpc/mysqlmanage/internal/svc"
	"awesomeProject/dou-yin/apps/commaction/rpc/mysqlmanage/types/mysqlmanageserver"
	"context"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type SubVideoFavoriteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSubVideoFavoriteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubVideoFavoriteLogic {
	return &SubVideoFavoriteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 取消点赞
func (l *SubVideoFavoriteLogic) SubVideoFavorite(in *mysqlmanageserver.SubVideoFavoriteRequest) (*mysqlmanageserver.SubVideoFavoriteResponse, error) {
	var n int64
	err := svc.DB.Table("video_info").Where("video_id = ?", in.VideoID).Count(&n).Error
	if err != nil {
		logx.Errorf("[pkg]logic [func]SubVideoFavorite [msg]gorm VideoID.Find [err]%v", err)
		return &mysqlmanageserver.SubVideoFavoriteResponse{Ok: 0}, nil
	}
	if int(n) == 0 {
		logx.Infof("[pkg]logic [func]AddVideoFavorite [msg]gorm Video does no exist")
		return &mysqlmanageserver.SubVideoFavoriteResponse{Ok: -2}, nil
	}

	err = svc.DB.Table("favorite_list").Where("favorite_video_id = ? AND favorite_user_id = ?", in.VideoID, in.UserID).Count(&n).Error
	if err != nil {
		logx.Errorf("[pkg]logic [func]SubVideoFavorite [msg]favorite_list.Count [err]%v", err)
		return &mysqlmanageserver.SubVideoFavoriteResponse{Ok: 0}, nil
	}
	if int(n) == 0 {
		logx.Infof("[pkg]logic [func]]SubVideoFavorite [msg]remove favorite already")
		return &mysqlmanageserver.SubVideoFavoriteResponse{Ok: -1}, nil
	}

	err = svc.DB.Table("video_info").Where("video_id = ?", in.VideoID).Update("favorite_count", gorm.Expr("favorite_count - 1")).Error
	if err != nil {
		logx.Errorf("[pkg]logic [func]SubVideoFavorite [msg]gorm favorite_count.UPDATE [err]%v", err)
		return &mysqlmanageserver.SubVideoFavoriteResponse{Ok: 0}, nil
	}

	err = svc.DB.Table("favorite_list").Where("favorite_video_id = ? AND favorite_user_id = ?", in.VideoID, in.UserID).Delete(Favorite_list{}).Error
	if err != nil {
		logx.Errorf("[pkg]logic [func]SubVideoFavorite [msg]gorm favorite_list.delete [err]%v", err)
		return &mysqlmanageserver.SubVideoFavoriteResponse{Ok: 0}, nil
	}
	return &mysqlmanageserver.SubVideoFavoriteResponse{Ok: 1}, err

}
