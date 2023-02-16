// Code generated by goctl. DO NOT EDIT.
// Source: video.proto

package server

import (
	"context"

	"awesomeProject/dou-yin/video/cmd/rpc/internal/logic"
	"awesomeProject/dou-yin/video/cmd/rpc/internal/svc"
	"awesomeProject/dou-yin/video/cmd/rpc/types/video"
)

type PublishVideoServer struct {
	svcCtx *svc.ServiceContext
	video.UnimplementedPublishVideoServer
}

func NewPublishVideoServer(svcCtx *svc.ServiceContext) *PublishVideoServer {
	return &PublishVideoServer{
		svcCtx: svcCtx,
	}
}

func (s *PublishVideoServer) PublishVideoAction(ctx context.Context, in *video.DouyinPublishActionRequest) (*video.DouyinPublishActionResponse, error) {
	l := logic.NewPublishVideoActionLogic(ctx, s.svcCtx)
	return l.PublishVideoAction(in)
}

func (s *PublishVideoServer) PublishVideoList(ctx context.Context, in *video.DouyinPublishListRequest) (*video.DouyinPublishListResponse, error) {
	l := logic.NewPublishVideoListLogic(ctx, s.svcCtx)
	return l.PublishVideoList(in)
}
