package video

import (
	"awesomeProject/dou-yin/apps/user/cmd/rpc/types/user"
	"awesomeProject/dou-yin/video/cmd/api/internal/svc"
	"awesomeProject/dou-yin/video/cmd/api/internal/types"
	"awesomeProject/dou-yin/video/cmd/rpc/types/video"
	"context"
	"sync"

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

	var group sync.WaitGroup
	group.Add(2)

	var info *user.DouyinUserResponse

	var res *video.DouyinPublishListResponse

	go func() {
		info, err = l.svcCtx.UserRpc.UserInfo(l.ctx, &user.DouyinUserRequest{
			UserId: req.UserId,
		})
		defer group.Done()
		if err != nil {
			logx.Error(err)
			return
		}
	}()

	go func() {
		res, err = l.svcCtx.VideoRpc.PublishVideoList(l.ctx, &video.DouyinPublishListRequest{
			UserId: req.UserId,
		})
		defer group.Done()
		if err != nil {
			return
		}
	}()

	group.Wait()

	videos := make([]types.Video, len(res.VideoList))

	for i, v := range res.VideoList {
		videos[i].Id = v.Id
		videos[i].Author = types.UserInfo{
			UserId:        info.UserId,
			UserName:      info.Username,
			FollowCount:   0,
			FollowerCount: 0,
			IsFollow:      false,
		}
		videos[i].PlayUrl = v.PlayUrl
		videos[i].CoverUrl = v.CoverUrl
		videos[i].FavoriteCount = 0
		videos[i].CommentCount = 0
		videos[i].IsFavorite = false
		videos[i].Title = v.Title

	}

	return &types.PublishVideoListResp{
		Videos: videos,
	}, nil
}
