package logic

import (
	"awesomeProject/dou-yin/apps/commaction/rpc/mysqlmanage/internal/svc"
	"awesomeProject/dou-yin/apps/commaction/rpc/mysqlmanage/types/mysqlmanageserver"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type IsFavotiteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIsFavotiteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsFavotiteLogic {
	return &IsFavotiteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 是否点赞
func (l *IsFavotiteLogic) IsFavotite(in *mysqlmanageserver.IsFavotiteRequest) (*mysqlmanageserver.IsFavotiteResponse, error) {

	var n int64
	err := svc.DB.Table("favorite_list").Where("favorite_video_id = ? and favorite_user_id = ?", in.VideoId, in.UserId).Count(&n).Error
	if err != nil {
		logx.Errorf("[pkg]logic [func]IsFavotite [msg]gorm VideoID = UserID  [err]%v", err)
		return &mysqlmanageserver.IsFavotiteResponse{Ok: false}, err
	}
	return &mysqlmanageserver.IsFavotiteResponse{Ok: n > 0}, nil
}
