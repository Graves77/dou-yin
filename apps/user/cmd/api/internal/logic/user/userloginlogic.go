package user

import (
	"awesomeProject/dou-yin/apps/user/cmd/api/internal/svc"
	"awesomeProject/dou-yin/apps/user/cmd/api/internal/types"
	"awesomeProject/dou-yin/apps/user/cmd/rpc/types/user"
	"awesomeProject/dou-yin/pkg/jwtx"
	"context"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	// todo: add your logic here and delete this line
	loginResp, err := l.svcCtx.UserRpc.Login(l.ctx, &user.DouyinUserLoginRequest{
		Username: req.Username,
		Password: req.Password,
	})

	if err != nil {
		return &types.LoginResponse{
			Status_code: 400,
			Status_msg:  "登录失败",
		}, nil
	}
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.JwtAuth.AccessExpire
	accessSecret := l.svcCtx.Config.JwtAuth.AccessSecret
	token, _ := jwtx.GetToken(accessSecret, now, accessExpire, loginResp.UserId)

	return &types.LoginResponse{
		Status_code: 200,
		Status_msg:  "登录成功",
		User_id:     loginResp.UserId,
		Token:       token,
	}, nil
}
