package user

import (
	"awesomeProject/dou-yin/apps/user/cmd/api/internal/svc"
	"awesomeProject/dou-yin/apps/user/cmd/api/internal/types"
	"awesomeProject/dou-yin/apps/user/cmd/rpc/types/user"
	"awesomeProject/dou-yin/common/jwtx"
	"context"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.RegisterRequest) (resp *types.RegisterResponse, err error) {
	// todo: add your logic here and delete this line
	response, err := l.svcCtx.UserRpc.Register(l.ctx, &user.DouyinUserRegisterRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return &types.RegisterResponse{
			Status_code: 400,
			Status_msg:  "注册失败",
		}, nil
	}
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.JwtAuth.AccessExpire
	accessSecret := l.svcCtx.Config.JwtAuth.AccessSecret
	token, _ := jwtx.GetToken(accessSecret, now, accessExpire, response.UserId)

	return &types.RegisterResponse{
		Status_code: 200,
		Status_msg:  "注册成功",
		User_id:     response.UserId,
		Token:       token,
	}, nil

}
