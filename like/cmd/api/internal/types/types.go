// Code generated by goctl. DO NOT EDIT.
package types

type LikeReq struct {
	VideoId    int64 `json:"videoId"`
	ActionType int64 `json:"actionType"` // true 点赞 / false 取消点赞
}

type LikeResp struct {
	StatusCode int64  `json:"statusCode"` // true 成功 / false 失败
	StatusMsg  string `json:"statusMsg"`
}

type LikeListReq struct {
	UserId int64 `json:"userId"`
}

type LikeListResp struct {
	Videos      []Video `json:"videos"`
	Status_code int32   `json:"status_code"`
	Status_msg  string  `json:"status_msg"`
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

type UserInfo struct {
	Id              int64  `json:"id"`
	Name            string `json:"name"`
	FollowCount     int64  `json:"followCount"`
	FollowerCount   int64  `json:"followerCount"`
	IsFollow        bool   `json:"isFollow"`
	Avatar          string `json:"avatar"`
	BackgroundImage string `json:"backgroundImage"`
	Signature       string `json:"signature"`
	TotalFavorited  string `json:"totalFavorited"`
	WorkCount       int64  `json:"workCount"`
	FavoriteCount   int64  `json:"favoriteCount"`
}