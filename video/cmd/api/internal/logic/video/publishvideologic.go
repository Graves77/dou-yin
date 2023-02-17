package video

import (
	"awesomeProject/dou-yin/video/cmd/api/internal/svc"
	"awesomeProject/dou-yin/video/cmd/api/internal/types"
	"awesomeProject/dou-yin/video/cmd/rpc/types/video"
	"context"
	"encoding/json"

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

	uid, _ := l.ctx.Value("uid").(json.Number).Int64()

	res, err := l.svcCtx.VideoRpc.PublishVideoAction(l.ctx, &video.DouyinPublishActionRequest{
		UserId:   uid,
		PlayUrl:  req.PlayUrl,
		CoverUrl: req.CoverUrl,
		Title:    req.Title,
	})
	if err != nil {
		return nil, err
	}

	return &types.PublishVideoActionResp{
		Status_code: 200,
		Status_msg:  "0",
		CoverUrl:    res.CoverUrl,
	}, nil
}
