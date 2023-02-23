package logic

import (
	"awesomeProject/dou-yin/apps/commaction/rpc/mysqlmanage/internal/svc"
	"awesomeProject/dou-yin/apps/commaction/rpc/mysqlmanage/types/mysqlmanageserver"
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

type RelationFollowerListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRelationFollowerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RelationFollowerListLogic {
	return &RelationFollowerListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 粉丝列表
func (l *RelationFollowerListLogic) RelationFollowerList(in *mysqlmanageserver.RelationFollowerListRequest) (*mysqlmanageserver.RelationFollowerListResponse, error) {
	testUserList_ := make([]*mysqlmanageserver.RelationUser, 0)

	db := svc.DB
	//#查找某个账号的粉丝列表   user_id：账号
	userInfoList, err := db.Raw(fmt.Sprintf("SELECT * FROM user_info where user_id in(SELECT follower_id FROM follow_and_follower_list where user_id=%d)", in.UserID)).Rows()

	if err != nil {
		return &mysqlmanageserver.RelationFollowerListResponse{}, err
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
	logx.Infof("\n\n%+v\n\n", &testUserList_)

	return &mysqlmanageserver.RelationFollowerListResponse{RelationUser: testUserList_}, nil
}
