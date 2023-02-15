package logic

import (
	"awesomeProject/dou-yin/video/cmd/rpc/internal/svc"
	"awesomeProject/dou-yin/video/cmd/rpc/types/video"
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
	// todo: add your logic here and delete this line

	return &video.DouyinPublishActionResponse{}, nil
}
