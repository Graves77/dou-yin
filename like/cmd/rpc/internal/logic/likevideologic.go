package logic

import (
	"awesomeProject/dou-yin/config"
	"awesomeProject/dou-yin/like/model"
	"context"
	"github.com/pkg/errors"

	"awesomeProject/dou-yin/like/cmd/rpc/internal/svc"
	"awesomeProject/dou-yin/like/cmd/rpc/types/like"

	"github.com/zeromicro/go-zero/core/logx"
)

type LikeVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLikeVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LikeVideoLogic {
	return &LikeVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LikeVideoLogic) LikeVideo(in *like.LikeVideoReq) (*like.LikeVideoResp, error) {
	likeData, err := l.svcCtx.FavoriteModel.FindOneByUserIdVideoId(l.ctx, in.UserId, in.VideoId)
	if errors.Is(err, model.ErrNotFound) {
		if in.StatusCode == config.UserLikeVideo {
			// 用户从未对此VideoId点赞过且点赞
			_, err := l.svcCtx.FavoriteModel.Insert(l.ctx, &model.Favorite{
				UserId:  in.UserId,
				VideoId: in.VideoId,
				Cancel:  config.UserLikeVideo,
			})
			if err != nil {
				return nil, err
			}
		}
		// 不存在且取消点赞，状态异常
		return &like.LikeVideoResp{
			StatusCode: false,
		}, nil
	}
	// 用户已经操作过，需要更新状态并且数据库中对状态和发送来对不同
	if likeData.Cancel != int64(in.StatusCode) {
		likeData.Cancel = int64(in.StatusCode)
		if err := l.svcCtx.FavoriteModel.Update(l.ctx, likeData); err != nil {
			return nil, err
		}
	}
	return &like.LikeVideoResp{
		StatusCode: true,
	}, nil
}
