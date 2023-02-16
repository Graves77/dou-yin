// Code generated by goctl. DO NOT EDIT.
package types

type PublishVideoActionReq struct {
	PlayUrl  string `json:"playUrl"`
	CoverUrl string `json:"coverUrl"`
	Title    string `json:"title"`
}

type PublishVideoActionResp struct {
	Status_code int32 `json:"status_code"`
	Status_msg  int32 `json:"status_msg"`
}

type PublishVideoListReq struct {
	UserId int64 `json:"userId"`
}

type PublishVideoListResp struct {
	Videos      []Video `json:"videos"`
	Status_code int32   `json:"status_code"`
	Status_msg  int32   `json:"status_msg"`
}

type UserInfo struct {
	UserId        int64  `json:"userId"`
	UserName      string `json:"userName"`
	FollowCount   int64  `json:"followCount"`
	FollowerCount int64  `json:"followerCount"`
	IsFollow      bool   `json:"isFollow"`
}

type Video struct {
	Id            int64    `json:"id"`
	Author        UserInfo `json:"userInfo"`
	PlayUrl       string   `json:"playUrl"`
	CoverUrl      string   `json:"coverUrl"`
	FavoriteCount int64    `json:"favoriteCount"`
	CommentCount  int64    `json:"commentCount"`
	IsFavorite    bool     `json:"isFavorite"`
	Title         string   `json:"title"`
}