package RelationFollowInterface

import (
	"awesomeProject/dou-yin/apps/commaction/cmd/rpc/mysqlmanage/types/mysqlmanageserver"
	"awesomeProject/dou-yin/apps/relationfollow/cmd/api/internal/svc"
	"awesomeProject/dou-yin/apps/relationfollow/cmd/api/internal/types"
	"awesomeProject/dou-yin/common/token"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type RelationActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRelationActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RelationActionLogic {
	return &RelationActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RelationActionLogic) RelationAction(req *types.RelationActionHandlerRequest) (resp *types.RelationActionHandlerResponse, err error) {
	ok, id, err := tools.CheckToke(req.Token)
	resultJson := types.RelationActionHandlerResponse{}

	if !ok {
		logx.Infof("[pkg]logic [func]PublishList [msg]req.Token is wrong ")
		return &types.RelationActionHandlerResponse{
			StatusCode: 400,
			StatusMsg:  "登录过期，请重新登陆",
		}, nil
	}
	if err != nil {
		logx.Errorf("[pkg]logic [func]PublishListr [msg]func CheckToken [err]%v", err)
		return &types.RelationActionHandlerResponse{
			StatusCode: 400,
			StatusMsg:  "Token校验出错",
		}, nil
	}

	result, err := l.svcCtx.MySQLManageRpc.RelationAction(l.ctx, &mysqlmanageserver.RelationActionRequest{
		UserID:     int64(id),
		ToUserID:   req.To_user_id,
		ActionType: int32(req.Sction_type),
	})
	if err != nil {
		resultJson.StatusCode = 402
		resultJson.StatusMsg = err.Error()
		return &resultJson, err
	}
	if result.Ok {
		resultJson.StatusCode = 0
		resultJson.StatusMsg = "success"
	} else {
		resultJson.StatusCode = 500
		resultJson.StatusMsg = err.Error()
	}
	return &resultJson, err
}
