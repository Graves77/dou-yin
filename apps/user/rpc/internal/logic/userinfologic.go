package logic

import (
	"context"

	"awesomeProject/dou-yin/apps/user/rpc/internal/svc"
	"awesomeProject/dou-yin/apps/user/rpc/types/user"

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

	return &user.DouyinUserResponse{}, nil
}
