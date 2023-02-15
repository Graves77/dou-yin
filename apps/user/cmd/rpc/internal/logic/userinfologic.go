package logic

import (
	"awesomeProject/dou-yin/apps/user/cmd/rpc/internal/svc"
	"awesomeProject/dou-yin/apps/user/cmd/rpc/types/user"
	"context"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserInfoLogic) UserInfo(in *user.DouyinUserRequest) (*user.DouyinUserResponse, error) {
	// todo: add your logic here and delete this line
	u, err := l.svcCtx.UserModel.FindOne(l.ctx, in.UserId)
	if err != nil {
		return nil, errors.New("查询不到该用户")
	}
	return &user.DouyinUserResponse{
		UserId:        u.UserId,
		Username:      u.Username,
		FollowCount:   u.FollowCount,
		FollowerCount: u.FollowerCount,
	}, nil
}
