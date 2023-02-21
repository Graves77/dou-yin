package CommActionInterface

import (
	"awesomeProject/dou-yin/apps/commaction/cmd/api/internal/svc"
	"awesomeProject/dou-yin/apps/commaction/cmd/api/internal/types"
	"awesomeProject/dou-yin/apps/commaction/cmd/rpc/mysqlmanage/types/mysqlmanageserver"
	"awesomeProject/dou-yin/common/minioconnect"
	"awesomeProject/dou-yin/common/token"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteListRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFavoriteListRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteListRegisterLogic {
	return &FavoriteListRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FavoriteListRegisterLogic) FavoriteListRegister(req *types.FavoriteListRegisterHandlerRequest) (resp *types.FavoriteListRegisterHandlerResponse, err error) {
	if req.Token == "" { //APP没有登录账号就请求了该API，为了不让前端报错，直接返回空数据
		return &types.FavoriteListRegisterHandlerResponse{
			StatusCode: 0,
		}, nil
	}
	ok, id, err := tools.CheckToke(req.Token)
	if !ok {
		logx.Infof("[pkg]logic [func]FavoriteListRegister [msg]req.Token is wrong ")
		return &types.FavoriteListRegisterHandlerResponse{
			StatusCode: 400,
			StatusMsg:  "登录过期，请重新登陆",
			VideoList:  []types.Video{},
		}, nil
	}
	if err != nil {
		logx.Errorf("[pkg]logic [func]FavoriteListRegister [msg]func CheckToken [err]%v", err)
		return &types.FavoriteListRegisterHandlerResponse{
			StatusCode: 400,
			StatusMsg:  "Token校验出错",
			VideoList:  []types.Video{},
		}, nil
	}

	n, err := l.svcCtx.MySQLManageRpc.FavoriteVideoNum(l.ctx, &mysqlmanageserver.FavoriteVideoNumRequest{
		UserID: int64(id),
	})
	if err != nil {
		logx.Errorf("[pkg]logic [func]FavoriteListRegister [msg]func FavoriteVideoNum [err]%v", err)
		return &types.FavoriteListRegisterHandlerResponse{
			StatusCode: 400,
			StatusMsg:  "操作失败",
			VideoList:  []types.Video{},
		}, nil
	}
	if n.N == -1 {
		logx.Infof("[pkg]logic [func]FavoriteListRegister [msg]User does not exit ")
		return &types.FavoriteListRegisterHandlerResponse{
			StatusCode: 400,
			StatusMsg:  "获取喜欢列表的用户信息失败",
			VideoList:  []types.Video{},
		}, nil
	}

	v, err := l.svcCtx.MySQLManageRpc.GetFavoriteVideoList(l.ctx, &mysqlmanageserver.GetFavoriteVideoListRequest{
		UserID: int64(id),
	})
	if err != nil {
		logx.Errorf("[pkg]logic [func]FavoriteListRegister [msg]func GetFavoriteVideoList [err]%v", err)
		return &types.FavoriteListRegisterHandlerResponse{
			StatusCode: 400,
			StatusMsg:  "操作失败",
			VideoList:  []types.Video{},
		}, nil
	}

	videolist := make([]types.Video, n.N)
	for i := 0; i < int(n.N); i++ {
		//user, ok := l.svcCtx.MySQLManageRpc.CheckUserInf(int(v[i].AuthorID), id)
		user, ok := l.svcCtx.MySQLManageRpc.CheckUserInf(l.ctx, &mysqlmanageserver.CheckUserInfRequest{
			UserId:     int64(v.VideoInfo[i].AuthorId),
			FollowerId: int64(id),
		})
		if ok != nil {
			logx.Infof("[pkg]logic [func]FavoriteListRegister [msg]User does not exist")
			return &types.FavoriteListRegisterHandlerResponse{
				StatusCode: 400,
				StatusMsg:  "获取视频对应的用户信息失败",
				VideoList:  []types.Video{},
			}, nil
		}
		realPlayUrl, _ := minioconnect.GetPlayUrl(v.VideoInfo[i].PlayUrl)
		realCoverUrl, _ := minioconnect.GetPlayUrl(v.VideoInfo[i].CoverUrl)
		videolist[i] = types.Video{
			Id: v.VideoInfo[i].VideoId,
			Author: types.User{
				UserId:        user.User.Users.UserId,
				Name:          user.User.Users.UserNickName,
				FollowCount:   user.User.Users.FollowCount,
				FollowerCount: user.User.Users.FollowerCount,
				IsFollow:      user.User.IsFollow,
			},
			PlayUrl:       realPlayUrl,
			CoverUrl:      realCoverUrl,
			FavoriteCount: v.VideoInfo[i].FavoriteCount,
			CommentCount:  v.VideoInfo[i].CommentCount,
			IsFavotite:    v.VideoInfo[i].IsFavotite,
			VideoTitle:    v.VideoInfo[i].VideoTitle,
		}
	}
	return &types.FavoriteListRegisterHandlerResponse{
		StatusCode: 0,
		StatusMsg:  "查询喜欢列表成功",
		VideoList:  videolist,
	}, err
}
