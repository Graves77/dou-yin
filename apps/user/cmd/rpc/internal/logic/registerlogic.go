package logic

import (
	"awesomeProject/dou-yin/apps/user/cmd/rpc/internal/svc"
	"awesomeProject/dou-yin/apps/user/cmd/rpc/types/user"
	"awesomeProject/dou-yin/apps/user/model"
	"awesomeProject/dou-yin/pkg/tool"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.DouyinUserRegisterRequest) (*user.DouyinUserRegisterResponse, error) {
	result, err := l.svcCtx.UserModel.Insert(l.ctx, &model.User{
		Username: in.Username,
		Password: tool.EnCoder(in.Password),
	})
	if err != nil {
		return nil, err
	}
	userId, _ := result.LastInsertId()
	return &user.DouyinUserRegisterResponse{
		UserId:   userId,
		Username: in.Username,
	}, nil
}
