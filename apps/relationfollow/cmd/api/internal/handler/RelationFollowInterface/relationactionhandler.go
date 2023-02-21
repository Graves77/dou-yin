package RelationFollowInterface

import (
	"awesomeProject/dou-yin/apps/relationfollow/cmd/api/internal/logic/RelationFollowInterface"
	"awesomeProject/dou-yin/apps/relationfollow/cmd/api/internal/svc"
	"awesomeProject/dou-yin/apps/relationfollow/cmd/api/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func RelationActionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RelationActionHandlerRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := RelationFollowInterface.NewRelationActionLogic(r.Context(), svcCtx)
		resp, err := l.RelationAction(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
