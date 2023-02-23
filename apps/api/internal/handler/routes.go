// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"awesomeProject/dou-yin/apps/api/internal/handler/CommActionInterface"
	RelationFollowInterface2 "awesomeProject/dou-yin/apps/api/internal/handler/RelationFollowInterface"
	"awesomeProject/dou-yin/apps/api/internal/handler/user"
	"awesomeProject/dou-yin/apps/api/internal/handler/video"
	"awesomeProject/dou-yin/apps/api/internal/svc"
	"github.com/dvwright/xss-mw"
	"github.com/gin-gonic/gin"

	"github.com/zeromicro/go-zero/rest"
	"net/http"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	router := gin.Default()
	var xssMdlwr xss.XssMw
	router.Use(xssMdlwr.RemoveXss())

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: user.UserLoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/register",
				Handler: user.UserRegisterHandler(serverCtx),
			},
		},
		rest.WithPrefix("/douyin/user"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/",
				Handler: user.UserInfoHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/douyin/user"),
	)
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/douyin/favorite/action",
				Handler: CommActionInterface.FavoriteActionHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/douyin/favorite/list",
				Handler: CommActionInterface.FavoriteListRegisterHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/douyin/comment/action",
				Handler: CommActionInterface.CommmentActionHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/douyin/comment/list",
				Handler: CommActionInterface.CommmentListHandler(serverCtx),
			},
		},
	)
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.UploadFile},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/action",
					Handler: video.PublishVideoHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/list/:userId",
					Handler: video.PublishVideoListHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),

		rest.WithPrefix("/douyin/publish"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/douyin/relation/action",
				Handler: RelationFollowInterface2.RelationActionHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/douyin/relation/follow/list",
				Handler: RelationFollowInterface2.RelationFollowListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/douyin/relation/follower/list",
				Handler: RelationFollowInterface2.RelationFollowerListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/douyin/relation/friend/list",
				Handler: RelationFollowInterface2.RelationFriendListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/douyin/message/chat",
				Handler: RelationFollowInterface2.MessageChatHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/douyin/message/action",
				Handler: RelationFollowInterface2.MessageActionHandler(serverCtx),
			},
		},
	)
}
