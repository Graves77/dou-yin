// Code generated by goctl. DO NOT EDIT.
// Source: video.proto

package server

import (
	logic2 "awesomeProject/dou-yin/apps/video/cmd/rpc/internal/logic"
	"awesomeProject/dou-yin/apps/video/cmd/rpc/internal/svc"
	video2 "awesomeProject/dou-yin/apps/video/cmd/rpc/types/video"
	"context"
)

type PublishVideoServer struct {
	svcCtx *svc.ServiceContext
	video2.UnimplementedPublishVideoServer
}

func NewPublishVideoServer(svcCtx *svc.ServiceContext) *PublishVideoServer {
	return &PublishVideoServer{
		svcCtx: svcCtx,
	}
}

func (s *PublishVideoServer) PublishVideoAction(ctx context.Context, in *video2.DouyinPublishActionRequest) (*video2.DouyinPublishActionResponse, error) {
	l := logic2.NewPublishVideoActionLogic(ctx, s.svcCtx)
	return l.PublishVideoAction(in)
}

func (s *PublishVideoServer) PublishVideoList(ctx context.Context, in *video2.DouyinPublishListRequest) (*video2.DouyinPublishListResponse, error) {
	l := logic2.NewPublishVideoListLogic(ctx, s.svcCtx)
	return l.PublishVideoList(in)
}
