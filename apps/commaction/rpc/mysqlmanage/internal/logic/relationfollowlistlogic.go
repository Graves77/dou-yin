package logic

import (
	"awesomeProject/dou-yin/apps/commaction/rpc/mysqlmanage/internal/svc"
	"awesomeProject/dou-yin/apps/commaction/rpc/mysqlmanage/types/mysqlmanageserver"
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

type RelationFollowListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRelationFollowListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RelationFollowListLogic {
	return &RelationFollowListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 关注列表
func (l *RelationFollowListLogic) RelationFollowList(in *mysqlmanageserver.RelationFollowListRequest) (*mysqlmanageserver.RelationFollowListResponse, error) {
	testUserList_ := make([]*mysqlmanageserver.RelationUser, 0)

	db := svc.DB
	//#查询某个账号关注的所有账号   follower_id：账号
	userInfoList, err := db.Raw(fmt.Sprintf("SELECT * FROM user_info where user_id in(SELECT user_id FROM follow_and_follower_list where follower_id = %d)", in.UserID)).Rows()
	if err != nil {
		return &mysqlmanageserver.RelationFollowListResponse{}, err
	} else {
		for userInfoList.Next() {
			tempUserList := mysqlmanageserver.RelationUser{}
			userInfoList.Scan(&tempUserList.Id, &tempUserList.Name, &tempUserList.FollowCount, &tempUserList.FollowerCount, &tempUserList.IsFollow) //！！err :查询出来的列数不同、数据格式不同时会报错，不影响格式正确的变量
			testUserList_ = append(testUserList_, &tempUserList)
		}
		//查询一遍上面查出来的id，是否已被当前登录的账号关注
		for i := 0; i < len(testUserList_); i++ {
			var ls CheckIsFollowLogic
			isFollow := mysqlmanageserver.CheckIsFollowRequest{UserId: int64(testUserList_[i].Id), FollowerId: int64(in.LoginUserID)}
			result, _ := ls.CheckIsFollow(&isFollow)
			if result.Ok {
				testUserList_[i].IsFollow = true
			}
		}
	}
	return &mysqlmanageserver.RelationFollowListResponse{RelationUser: testUserList_}, nil

}
