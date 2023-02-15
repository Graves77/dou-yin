package video

import (
	"awesomeProject/dou-yin/video/cmd/api/internal/svc"
	"awesomeProject/dou-yin/video/cmd/api/internal/types"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishVideoListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPublishVideoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishVideoListLogic {
	return &PublishVideoListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublishVideoListLogic) PublishVideoList(req *types.PublishVideoListReq) (resp *types.PublishVideoListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
