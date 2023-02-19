// Code generated by goctl. DO NOT EDIT.
// Source: like.proto

package server

import (
	"context"

	"awesomeProject/dou-yin/like/cmd/rpc/internal/logic"
	"awesomeProject/dou-yin/like/cmd/rpc/internal/svc"
	"awesomeProject/dou-yin/like/cmd/rpc/types/like"
)

type LikeServer struct {
	svcCtx *svc.ServiceContext
	like.UnimplementedLikeServer
}

func NewLikeServer(svcCtx *svc.ServiceContext) *LikeServer {
	return &LikeServer{
		svcCtx: svcCtx,
	}
}

func (s *LikeServer) LikeVideo(ctx context.Context, in *like.LikeVideoReq) (*like.LikeVideoResp, error) {
	l := logic.NewLikeVideoLogic(ctx, s.svcCtx)
	return l.LikeVideo(in)
}

func (s *LikeServer) LikeList(ctx context.Context, in *like.LikeListReq) (*like.LikeListResp, error) {
	l := logic.NewLikeListLogic(ctx, s.svcCtx)
	return l.LikeList(in)
}