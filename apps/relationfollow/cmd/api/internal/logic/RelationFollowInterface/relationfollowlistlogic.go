package RelationFollowInterface

import (
	"awesomeProject/dou-yin/apps/commaction/cmd/rpc/mysqlmanage/types/mysqlmanageserver"
	"awesomeProject/dou-yin/apps/relationfollow/cmd/api/internal/svc"
	"awesomeProject/dou-yin/apps/relationfollow/cmd/api/internal/types"
	"awesomeProject/dou-yin/common/token"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type RelationFollowListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRelationFollowListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RelationFollowListLogic {
	return &RelationFollowListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RelationFollowListLogic) RelationFollowList(req *types.RelationFollowListHandlerRequest) (resp *types.RelationFollowListHandlerResponse, err error) {
	ok, id, err := tools.CheckToke(req.Token)
	if !ok {
		logx.Infof("[pkg]logic [func]PublishList [msg]req.Token is wrong ")
		return &types.RelationFollowListHandlerResponse{
			StatusCode: 400,
			StatusMsg:  "登录过期，请重新登陆",
		}, nil
	}
	if err != nil {
		logx.Errorf("[pkg]logic [func]PublishListr [msg]func CheckToken [err]%v", err)
		return &types.RelationFollowListHandlerResponse{
			StatusCode: 400,
			StatusMsg:  "Token校验出错",
		}, nil
	}
	resultJson := &types.RelationFollowListHandlerResponse{}
	rflhr, err := l.svcCtx.MySQLManageRpc.RelationFollowList(l.ctx, &mysqlmanageserver.RelationFollowListRequest{
		LoginUserID: int64(id),
		UserID:      req.UserId,
	})
	if err != nil {
		resultJson.StatusCode = 400
		resultJson.StatusMsg = err.Error()
		return resultJson, err
	}
	Userlist := make([]types.RelationUser, 0)
	for i := 0; i < len(rflhr.RelationUser); i++ {
		Userlist = append(Userlist, types.RelationUser{
			Id:            rflhr.RelationUser[i].Id,
			Name:          rflhr.RelationUser[i].Name,
			FollowCount:   rflhr.RelationUser[i].FollowCount,
			FollowerCount: rflhr.RelationUser[i].FollowerCount,
			IsFollow:      rflhr.RelationUser[i].IsFollow,
		})
		resultJson.StatusCode = 0
		resultJson.StatusMsg = "success"
	}
	resultJson.UserList = Userlist
	return resultJson, err
}
