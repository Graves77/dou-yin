package CommActionInterface

import (
	"awesomeProject/dou-yin/apps/api/internal/svc"
	"awesomeProject/dou-yin/apps/api/internal/types"
	"awesomeProject/dou-yin/apps/commaction/rpc/mysqlmanage/types/mysqlmanageserver"
	tools "awesomeProject/dou-yin/common/token"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFavoriteActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteActionLogic {
	return &FavoriteActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FavoriteActionLogic) FavoriteAction(req *types.FavoriteActionHandlerRequest) (resp *types.FavoriteActionHandlerResponse, err error) {
	ok, id, err := tools.CheckToke(req.Token)
	if !ok {
		logx.Infof("[pkg]logic [func]FavoriteAction [msg]req.Token is wrong ")
		return &types.FavoriteActionHandlerResponse{
			StatusCode: 402,
			StatusMsg:  "登录过期，请重新登陆",
		}, nil
	}
	if err != nil {
		logx.Errorf("[pkg]logic [func]FavoriteAction [msg]func CheckToken [err]%v", err)
		return &types.FavoriteActionHandlerResponse{
			StatusCode: 400,
			StatusMsg:  "Token校验出错",
		}, nil
	}

	if req.ActionType == 1 {
		//ok, err := l.svcCtx.MySQLManageRpc.AddVideoFavorite(int64(id), req.VideoId)
		ok, err := l.svcCtx.MySQLManageRpc.AddVideoFavorite(l.ctx, &mysqlmanageserver.AddVideoFavoriteRequest{
			UserID:  int64(id),
			VideoID: req.VideoId,
		})
		if ok.Ok == 0 || err != nil {
			logx.Errorf("[pkg]logic [func]FavoriteAction [msg]func AddVideoFavorite [err]%v", err)
			return &types.FavoriteActionHandlerResponse{
				StatusCode: 400,
				StatusMsg:  "点赞失败，稍后重试",
			}, nil
		}
		if ok.Ok == -1 {
			logx.Infof("[pkg]logic [func]FavoriteAction [msg]favorite already")
			return &types.FavoriteActionHandlerResponse{
				StatusCode: 400,
				StatusMsg:  "无法点赞不存在的视频",
			}, nil
		}
		if ok.Ok == -2 {
			logx.Infof("[pkg]logic [func]FavoriteAction [msg]Video does no exist")
			return &types.FavoriteActionHandlerResponse{
				StatusCode: 400,
				StatusMsg:  "已经点赞过",
			}, nil
		}
		return &types.FavoriteActionHandlerResponse{
			StatusCode: 0,
			StatusMsg:  "点赞成功",
		}, nil
	} else if req.ActionType == 2 {
		//ok, err := l.svcCtx.MySQLManageRpc.SubVideoFavorite(int64(id), req.VideoId)
		ok, err := l.svcCtx.MySQLManageRpc.SubVideoFavorite(l.ctx, &mysqlmanageserver.SubVideoFavoriteRequest{
			UserID:  int64(id),
			VideoID: req.VideoId,
		})
		if ok.Ok == 0 || err != nil {
			logx.Errorf("[pkg]logic [func]FavoriteAction [msg]func SubVideoFavorite [err]%v", err)
			return &types.FavoriteActionHandlerResponse{
				StatusCode: 400,
				StatusMsg:  "取消点赞失败，稍后重试",
			}, nil
		}
		if ok.Ok == -1 {
			logx.Infof("[pkg]logic [func]FavoriteAction [msg]Video does no exist")
			return &types.FavoriteActionHandlerResponse{
				StatusCode: 400,
				StatusMsg:  "已经取消点赞过",
			}, nil
		}
		if ok.Ok == -2 {
			logx.Infof("[pkg]logic [func]FavoriteAction [msg]favorite already")
			return &types.FavoriteActionHandlerResponse{
				StatusCode: 402,
				StatusMsg:  "无法取消点赞不存在的视频",
			}, nil
		}
		return &types.FavoriteActionHandlerResponse{
			StatusCode: 0,
			StatusMsg:  "取消点赞成功",
		}, nil
	}

	return &types.FavoriteActionHandlerResponse{
		StatusCode: 400,
		StatusMsg:  "操作状态异常",
	}, nil

}
