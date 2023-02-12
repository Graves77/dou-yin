package logic

import (
	"context"

	"awesomeProject/dou-yin/apps/user/rpc/internal/svc"
	"awesomeProject/dou-yin/apps/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.DouyinUserLoginRequest) (*user.DouyinUserLoginResponse, error) {
	// todo: add your logic here and delete this line

	return &user.DouyinUserLoginResponse{}, nil
}
