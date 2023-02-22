package RelationFollowInterface

import (
	"awesomeProject/dou-yin/apps/api/internal/svc"
	"awesomeProject/dou-yin/apps/api/internal/types"
	"awesomeProject/dou-yin/apps/commaction/rpc/mongodbmanage/types/mongodbmanageserver"

	"awesomeProject/dou-yin/common/token"

	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type MessageChatLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMessageChatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MessageChatLogic {
	return &MessageChatLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MessageChatLogic) MessageChat(req *types.MessageChatHandlerRequest) (resp *types.MessageChatHandlerResponse, err error) {
	ok, uid, err := tools.CheckToke(req.Token)
	if err != nil {
		logx.Errorf("[pkg]RelationFollowInterface [func]MessageChat [msg]checktoken failed, [err]%v", err)
		return &types.MessageChatHandlerResponse{
			StatusCode:  400,
			StatusMsg:   "checktoken failed",
			MessageList: []types.SingleMessage{},
		}, err
	}
	if !ok {
		logx.Errorf("[pkg]RelationFollowInterface [func]MessageChat [msg]请重新登录, [err]%v", err)
		return &types.MessageChatHandlerResponse{
			StatusCode:  400,
			StatusMsg:   "请重新登录",
			MessageList: []types.SingleMessage{},
		}, err
	}

	messageList, err := l.svcCtx.MongoDBMangerRpc.GetMessage(l.ctx, &mongodbmanageserver.MessageChatRequest{
		ToUserId:   req.ToUserId,
		FromUserId: int64(uid),
	})
	if err != nil {
		logx.Errorf("[pkg]RelationFollowInterface [func]MessageChat [msg]get message list failed, [err]%v", err)
		return &types.MessageChatHandlerResponse{
			StatusCode:  400,
			StatusMsg:   "get message list failed",
			MessageList: []types.SingleMessage{},
		}, err
	}

	var res []types.SingleMessage
	for _, v := range messageList.MessageList {
		message := types.SingleMessage{
			Id:         v.Id,
			Content:    v.Content,
			CreateTime: v.CreateTime,
		}
		res = append(res, message)
	}
	return &types.MessageChatHandlerResponse{
		StatusCode:  2000,
		StatusMsg:   "get message list success",
		MessageList: res,
	}, nil
}
