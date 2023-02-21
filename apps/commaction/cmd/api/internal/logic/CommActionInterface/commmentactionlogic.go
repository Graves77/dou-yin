package CommActionInterface

import (
	"awesomeProject/dou-yin/apps/commaction/cmd/api/internal/svc"
	"awesomeProject/dou-yin/apps/commaction/cmd/api/internal/types"
	"awesomeProject/dou-yin/apps/commaction/cmd/rpc/mongodbmanage/types/mongodbmanageserver"
	"awesomeProject/dou-yin/common/token"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommmentActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommmentActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommmentActionLogic {
	return &CommmentActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommmentActionLogic) CommmentAction(req *types.CommmentActionHandlerRequest) (resp *types.CommmentActionHandlerResponse, err error) {
	//parse token
	flag, userId, err := tools.CheckToke(req.Token)
	if !flag || err != nil {
		logx.Errorf("[pkg]logic [func]CommentAction [msg]parse token failed, [err]%v", err)
		return &types.CommmentActionHandlerResponse{
			StatusCode: 400,
			StatusMsg:  "parse token failed",
		}, err
	}
	rpcResponse, err := l.svcCtx.MongoDBMangerRpc.MakeComment(l.ctx, &mongodbmanageserver.CommentActionRequest{
		UserId:      int64(userId),
		VideoId:     req.VideoId,
		ActionType:  req.ActionType,
		CommentText: &req.CommentText,
		CommentId:   &req.CommentId,
	})
	if err != nil {
		logx.Errorf("[pkg]logic [func]CommentAction [msg]make comment failed, [err]%v", err)
		return &types.CommmentActionHandlerResponse{
			StatusCode: 400,
			StatusMsg:  "make comment failed",
			Comment:    types.CommentResp{},
		}, err
	}

	if req.ActionType == 2 {
		return &types.CommmentActionHandlerResponse{
			StatusCode: 200,
			StatusMsg:  "delete comment success",
			Comment:    types.CommentResp{},
		}, nil
	}

	comment := types.CommentResp{
		Id: rpcResponse.Comment.Id,
		User: types.User{
			UserId:        rpcResponse.Comment.User.UserId,
			Name:          rpcResponse.Comment.User.UserNickName,
			FollowCount:   rpcResponse.Comment.User.FollowCount,
			FollowerCount: rpcResponse.Comment.User.FollowerCount,
			IsFollow:      rpcResponse.Comment.User.IsFollow,
		},
		Content:    rpcResponse.Comment.Content,
		CreateDate: rpcResponse.Comment.CreateDate,
	}

	return &types.CommmentActionHandlerResponse{
		StatusCode: 0,
		StatusMsg:  "make comment success",
		Comment:    comment,
	}, nil
}
