package logic

import (
	"awesomeProject/dou-yin/apps/commaction/rpc/mongodbmanage/internal/svc"
	"awesomeProject/dou-yin/apps/commaction/rpc/mongodbmanage/types/mongodbmanageserver"
	"awesomeProject/dou-yin/common/mongodb"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentLogic {
	return &GetCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCommentLogic) GetComment(in *mongodbmanageserver.CommentListRequest) (*mongodbmanageserver.CommentListResponse, error) {
	var commentList []*mongodbmanageserver.Comment
	videoId := in.VideoId
	collection := mongodb.MongoDatabase.Collection("comment")
	filter := bson.D{{
		Key:   "video_id",
		Value: videoId,
	}}
	opts := &options.FindOptions{}
	sortOption := bson.D{{
		Key:   "_id",
		Value: -1,
	}}
	opts.Sort = sortOption
	cur, err := collection.Find(context.Background(), filter, opts)
	if err != nil {
		logx.Errorf("[pkg]logic [func]GetComment [msg]find comments failed, [err]%v", err)
		return &mongodbmanageserver.CommentListResponse{
			CommentList: nil,
		}, err
	}

	for cur.Next(context.Background()) {
		var comment commentForMongo
		err = cur.Decode(&comment)
		if err != nil {
			logx.Errorf("[pkg]logic [func]GetComment [msg]decode comment failed, [err]%v", err)
			return &mongodbmanageserver.CommentListResponse{
				CommentList: nil,
			}, err
		}
		temp := &mongodbmanageserver.Comment{
			Id: comment.Id,
			User: &mongodbmanageserver.User{
				UserId:        comment.User.UserId,
				UserNickName:  comment.User.Name,
				FollowCount:   comment.User.FollowCount,
				FollowerCount: comment.User.FollowerCount,
				IsFollow:      comment.User.IsFollow,
			},
			Content:    comment.Content,
			CreateDate: comment.CreateDate,
		}
		commentList = append(commentList, temp)
	}
	err = cur.Err()
	if err != nil {
		logx.Errorf("[pkg]logic [func]GetComment [msg]cur has an error, [err]%v", err)
		return &mongodbmanageserver.CommentListResponse{
			CommentList: nil,
		}, err
	}
	cur.Close(context.Background())
	return &mongodbmanageserver.CommentListResponse{
		CommentList: commentList,
	}, nil
}
