package RelationFollowInterface

import (
	"awesomeProject/dou-yin/apps/api/internal/logic/RelationFollowInterface"
	"awesomeProject/dou-yin/apps/api/internal/svc"
	"awesomeProject/dou-yin/apps/api/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func RelationFriendListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RelationFriendListHandlerRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := RelationFollowInterface.NewRelationFriendListLogic(r.Context(), svcCtx)
		resp, err := l.RelationFriendList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
