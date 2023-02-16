package video

import (
	"awesomeProject/dou-yin/video/cmd/api/internal/svc"
	"awesomeProject/dou-yin/video/cmd/api/internal/types"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishVideoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPublishVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishVideoLogic {
	return &PublishVideoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublishVideoLogic) PublishVideo(req *types.PublishVideoActionReq) (resp *types.PublishVideoActionResp, err error) {
	// todo: add your logic here and delete this line

	return
}