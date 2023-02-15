package logic

import (
	"awesomeProject/dou-yin/apps/user/cmd/rpc/internal/svc"
	"awesomeProject/dou-yin/apps/user/cmd/rpc/types/user"
	"awesomeProject/dou-yin/pkg/tool"
	"context"
	"errors"

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
	u, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, in.Username)
	if err != nil {
		return nil, err
	}
	password := tool.EnCoder(in.Password)
	if u.Password != password {
		return nil, errors.New("用户名或密码错误")
	}
	return &user.DouyinUserLoginResponse{
		UserId:   u.UserId,
		Username: u.Username,
	}, nil
}
