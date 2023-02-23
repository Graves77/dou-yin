package logic

import (
	"awesomeProject/dou-yin/apps/commaction/rpc/mysqlmanage/internal/svc"
	"awesomeProject/dou-yin/apps/commaction/rpc/mysqlmanage/types/mysqlmanageserver"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteVideoNumLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFavoriteVideoNumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteVideoNumLogic {
	return &FavoriteVideoNumLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 喜欢视频数量
func (l *FavoriteVideoNumLogic) FavoriteVideoNum(in *mysqlmanageserver.FavoriteVideoNumRequest) (*mysqlmanageserver.FavoriteVideoNumResponse, error) {
	var exist int64
	var n int64

	err := svc.DB.Table("user_info").Where("user_id = ?", in.UserID).Count(&exist).Error
	if err != nil {
		logx.Errorf("[pkg]logic [func]FavoriteVideoNum [msg]gorm UserID.Count [err]%v", err)
		return &mysqlmanageserver.FavoriteVideoNumResponse{N: 0}, nil
	}
	if exist == 0 {
		logx.Infof("[pkg]logic [func]FavoriteVideoNum [msg]User does not exit")
		return &mysqlmanageserver.FavoriteVideoNumResponse{N: -1}, err
	}

	err = svc.DB.Table("favorite_list").Where("favorite_user_id = ?", in.UserID).Count(&n).Error
	if err != nil {
		logx.Errorf("[pkg]logic [func]FavoriteVideoNum [msg]gorm avorite_user_id.Count [err]%v", err)
		return &mysqlmanageserver.FavoriteVideoNumResponse{N: n}, nil
	}
	return &mysqlmanageserver.FavoriteVideoNumResponse{N: n}, err

}
