package like

import (
	"context"

	"awesomeProject/dou-yin/like/cmd/api/internal/svc"
	"awesomeProject/dou-yin/like/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LikeListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLikeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LikeListLogic {
	return &LikeListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LikeListLogic) LikeList(req *types.LikeListReq) (resp *types.LikeListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
