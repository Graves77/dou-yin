package RelationFollowInterface

import (
	"awesomeProject/dou-yin/apps/commaction/cmd/rpc/mongodbmanage/types/mongodbmanageserver"
	"awesomeProject/dou-yin/apps/relationfollow/cmd/api/internal/svc"
	"awesomeProject/dou-yin/apps/relationfollow/cmd/api/internal/types"
	"awesomeProject/dou-yin/common/token"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type MessageActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMessageActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MessageActionLogic {
	return &MessageActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MessageActionLogic) MessageAction(req *types.MessageActionHandlerRequest) (resp *types.MessageActionHandlerResponse, err error) {
	ok, uid, err := tools.CheckToke(req.Token)
	if err!=nil {
		logx.Errorf("[pkg]RelationFollowInterface [func]MessageAction [msg]checktoken failed, [err]%v", err)
		return &types.MessageActionHandlerResponse{
			StatusCode: 400,
			StatusMsg: "checktoken failed",
		}, err
	}
	if !ok {
		logx.Errorf("[pkg]RelationFollowInterface [func]MessageAction [msg]请重新登录, [err]%v", err)
		return &types.MessageActionHandlerResponse{
			StatusCode: 400,
			StatusMsg: "请重新登录",
		}, err
	}

	_, err = l.svcCtx.MongoDBMangerRpc.SendMessage(l.ctx, &mongodbmanageserver.MessageActionRequest{
		ToUserId: req.ToUserId,
		ActionType: req.ActionType,
		Content: req.Content,
		FromUserId: int64(uid),
	})
	if err!=nil {
		logx.Errorf("[pkg]RelationFollowInterface [func]MessageAction [msg]get message list failed, [err]%v", err)
		return &types.MessageActionHandlerResponse{
			StatusCode: 400,
			StatusMsg: "send message failed",
		}, err
	}
	return &types.MessageActionHandlerResponse{
		StatusCode: 200,
		StatusMsg: "send message success",
	}, nil
}
