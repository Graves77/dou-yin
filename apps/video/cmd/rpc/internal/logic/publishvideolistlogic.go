package logic

import (
	"awesomeProject/dou-yin/apps/video/cmd/rpc/internal/svc"
	"awesomeProject/dou-yin/apps/video/cmd/rpc/types/video"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishVideoListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPublishVideoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishVideoListLogic {
	return &PublishVideoListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PublishVideoListLogic) PublishVideoList(in *video.DouyinPublishListRequest) (*video.DouyinPublishListResponse, error) {
	videos, err := l.svcCtx.VideoModel.FindListByUid(l.ctx, in.UserId)
	if err != nil {
		return nil, err
	}
	videoresp := make([]*video.Video, len(videos))

	for i, v := range videos {
		videoresp[i] = &video.Video{
			Id:            v.VideoId,
			UserId:        v.UserId,
			PlayUrl:       v.PlayUrl,
			CoverUrl:      v.CoverUrl,
			FavoriteCount: v.FavoriteCount,
			CommentCount:  v.CommentCount,
			IsFavorite:    false,
			Title:         v.Title,
		}
	}

	return &video.DouyinPublishListResponse{
		VideoList: videoresp,
	}, nil
}
