package logic

import (
	"awesomeProject/dou-yin/apps/video/cmd/rpc/internal/svc"
	"awesomeProject/dou-yin/apps/video/cmd/rpc/types/video"
	"awesomeProject/dou-yin/apps/video/model"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type PublishVideoActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPublishVideoActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishVideoActionLogic {
	return &PublishVideoActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PublishVideoActionLogic) PublishVideoAction(in *video.DouyinPublishActionRequest) (*video.DouyinPublishActionResponse, error) {
	v, err := l.svcCtx.VideoModel.Insert(l.ctx, &model.Videos{
		UserId:   in.UserId,
		PlayUrl:  in.PlayUrl,
		CoverUrl: in.CoverUrl,
		Title:    in.Title,
	})
	if err != nil {
		return nil, err
	}
	id, _ := v.LastInsertId()
	return &video.DouyinPublishActionResponse{
		UserId:   id,
		PlayUrl:  in.PlayUrl,
		CoverUrl: in.CoverUrl,
		Title:    in.Title,
	}, nil
}
