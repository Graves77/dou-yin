package CommActionInterface

import (
	"awesomeProject/dou-yin/apps/commaction/cmd/api/internal/svc"
	"awesomeProject/dou-yin/apps/commaction/cmd/api/internal/types"
	"awesomeProject/dou-yin/apps/commaction/cmd/rpc/mongodbmanage/types/mongodbmanageserver"
	"awesomeProject/dou-yin/common/token"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommmentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommmentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommmentListLogic {
	return &CommmentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommmentListLogic) CommmentList(req *types.CommmentListHandlerRequest) (resp *types.CommmentListHandlerResponse, err error) {
	flag, _, err := tools.CheckToke(req.Token)
	if !flag || err != nil {
		logx.Errorf("[pkg]logic [func]CommmentList [msg]parse token failed, [err]%v", err)
		return &types.CommmentListHandlerResponse{
			StatusCode: 400,
			StatusMsg:  "parse token failed",
		}, err
	}

	rpcResponse, err := l.svcCtx.MongoDBMangerRpc.GetComment(l.ctx, &mongodbmanageserver.CommentListRequest{
		VideoId: req.VideoId,
	})
	if err != nil {
		logx.Errorf("[pkg]logic [func]CommmentList [msg]get comment list, [err]%v", err)
		return &types.CommmentListHandlerResponse{
			StatusCode: 400,
			StatusMsg:  "get comment list",
		}, err
	}
	var commentList []types.CommentResp
	for _, v := range rpcResponse.CommentList {
		singleComment := types.CommentResp{
			Id: v.Id,
			User: types.User{
				UserId:        v.User.UserId,
				Name:          v.User.UserNickName,
				FollowCount:   v.User.FollowCount,
				FollowerCount: v.User.FollowerCount,
				IsFollow:      v.User.IsFollow,
			},
			Content:    v.Content,
			CreateDate: v.CreateDate,
		}
		commentList = append(commentList, singleComment)
	}
	return &types.CommmentListHandlerResponse{
		StatusCode:  0,
		StatusMsg:   "get message list success",
		CommentList: commentList,
	}, nil
}
