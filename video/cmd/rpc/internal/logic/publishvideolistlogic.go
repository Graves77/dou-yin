package logic

import (
	"awesomeProject/dou-yin/video/cmd/rpc/internal/svc"
	"awesomeProject/dou-yin/video/cmd/rpc/types/video"
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
	// todo: add your logic here and delete this line

	return &video.DouyinPublishListResponse{}, nil
}
