package user

import (
	"awesomeProject/dou-yin/apps/user/cmd/api/internal/svc"
	"awesomeProject/dou-yin/apps/user/cmd/api/internal/types"
	"awesomeProject/dou-yin/apps/user/cmd/rpc/types/user"
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo() (resp *types.UserInfoResponse, err error) {
	// todo: add your logic here and delete this line
	value := l.ctx.Value("uid")
	uid, _ := value.(json.Number).Int64()
	info, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.DouyinUserRequest{
		UserId: uid,
	})
	if err != nil {
		return &types.UserInfoResponse{
			Status_code: 400,
			Status_msg:  "用户不存在",
		}, nil
	}
	return &types.UserInfoResponse{
		Status_code:    200,
		Status_msg:     "success",
		Id:             uid,
		Name:           info.GetUsername(),
		Follow_count:   info.GetFollowCount(),
		Follower_count: info.GetFollowerCount(),
		Is_follow:      false,
	}, nil
}
