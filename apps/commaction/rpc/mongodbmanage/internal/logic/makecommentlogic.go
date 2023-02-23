package logic

import (
	"awesomeProject/dou-yin/apps/commaction/rpc/mongodbmanage/internal/svc"
	"awesomeProject/dou-yin/apps/commaction/rpc/mongodbmanage/types/mongodbmanageserver"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type userForMongo struct {
	UserId        int64  `gorm:"column:user_id"        json:"id"               form:"user_id"        bson:"user_id"`
	Name          string `gorm:"column:user_nick_name" json:"name"             form:"name"           bson:"name"`
	FollowCount   int64  `gorm:"column:follow_count"   json:"follow_count"     form:"follow_count"   bson:"follow_count"`
	FollowerCount int64  `gorm:"column:follower_count" json:"follower_count"   form:"follower_count" bson:"follower_count"`
	IsFollow      bool   `json:"is_follow"             form:"is_follow"        bson:"is_follow"`
}

type commentForMongo struct {
	Id         int64        `json:"id"          form:"id"         bson:"_id"`
	VideoId    int64        `json:"video_id"    form:"video_id"    bson:"video_id"` //视频id
	User       userForMongo `json:"user"        form:"user"        bson:"user"`
	Content    string       `json:"content"     form:"content"     bson:"content"`
	CreateDate string       `json:"create_date" form:"create_date" bson:"create_date"`
}
type MakeCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMakeCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MakeCommentLogic {
	return &MakeCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MakeCommentLogic) MakeComment(in *mongodbmanageserver.CommentActionRequest) (*mongodbmanageserver.CommentActionResponse, error) {

	return &mongodbmanageserver.CommentActionResponse{}, nil
}
