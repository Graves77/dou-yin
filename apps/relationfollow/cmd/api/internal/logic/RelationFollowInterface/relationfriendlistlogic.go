package RelationFollowInterface

import (
	"awesomeProject/dou-yin/apps/commaction/cmd/rpc/mysqlmanage/types/mysqlmanageserver"
	"awesomeProject/dou-yin/apps/relationfollow/cmd/api/internal/svc"
	"awesomeProject/dou-yin/apps/relationfollow/cmd/api/internal/types"
	"awesomeProject/dou-yin/common/token"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type RelationFriendListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRelationFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RelationFriendListLogic {
	return &RelationFriendListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RelationFriendListLogic) RelationFriendList(req *types.RelationFriendListHandlerRequest) (resp *types.RelationFriendListHandlerResponse, err error) {

	ok, id, err := tools.CheckToke(req.Token)
	if !ok {
		logx.Infof("[pkg]logic [func]PublishList [msg]req.Token is wrong ")
		return &types.RelationFriendListHandlerResponse{
			StatusCode: 400,
			StatusMsg:  "登录过期，请重新登陆",
		}, nil
	}
	if err != nil {
		logx.Errorf("[pkg]logic [func]PublishListr [msg]func CheckToken [err]%v", err)
		return &types.RelationFriendListHandlerResponse{
			StatusCode: 400,
			StatusMsg:  "Token校验出错",
		}, nil
	}
	resultJson := &types.RelationFriendListHandlerResponse{}
	rflhr, err := l.svcCtx.MySQLManageRpc.RelationFollowList(l.ctx, &mysqlmanageserver.RelationFollowListRequest{
		LoginUserID: int64(id),
		UserID:      req.UserId,
	})
	if err != nil {
		resultJson.StatusCode = 402
		resultJson.StatusMsg = err.Error()
		return resultJson, err
	}
	Userlist := make([]types.FriendUser, 0)
	for i := 0; i < len(rflhr.RelationUser); i++ {
		Userlist = append(Userlist, types.FriendUser{
			Id:            rflhr.RelationUser[i].Id,
			Name:          rflhr.RelationUser[i].Name,
			FollowCount:   rflhr.RelationUser[i].FollowCount,
			FollowerCount: rflhr.RelationUser[i].FollowerCount,
			IsFollow:      rflhr.RelationUser[i].IsFollow,
			Avatar:        "http://www.bailinzhe.com/image/2019-11-08/0ab5979f578d4f8bf15b20e6e51f0f2a.jpg",
			Message:       "msg",
			MsgType:       1,
		})
		resultJson.StatusCode = 0
		resultJson.StatusMsg = "success"
	}
	resultJson.Userlist = Userlist
	return resultJson, err
}
