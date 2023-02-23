package logic

import (
	"awesomeProject/dou-yin/apps/commaction/rpc/mysqlmanage/internal/svc"
	"awesomeProject/dou-yin/apps/commaction/rpc/mysqlmanage/types/mysqlmanageserver"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type VideoNumLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewVideoNumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VideoNumLogic {
	return &VideoNumLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 发布列表视频数量
func (l *VideoNumLogic) VideoNum(in *mysqlmanageserver.VideoNumRequest) (*mysqlmanageserver.VideoNumResponse, error) {
	var num int64
	err := svc.DB.Table("video_info").Where("author_id = ?", in.AuthorId).Count(&num).Error
	if err != nil {
		logx.Errorf("[pkg]logic [func]VideoNum [msg]gorm video_info.Count [err]%v", err)
		return &mysqlmanageserver.VideoNumResponse{}, err
	}
	return &mysqlmanageserver.VideoNumResponse{Num: num}, nil
}
