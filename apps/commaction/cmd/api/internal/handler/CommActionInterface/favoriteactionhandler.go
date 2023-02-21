package CommActionInterface

import (
	"awesomeProject/dou-yin/apps/commaction/cmd/api/internal/logic/CommActionInterface"
	"awesomeProject/dou-yin/apps/commaction/cmd/api/internal/svc"
	"awesomeProject/dou-yin/apps/commaction/cmd/api/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func FavoriteActionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FavoriteActionHandlerRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := CommActionInterface.NewFavoriteActionLogic(r.Context(), svcCtx)
		resp, err := l.FavoriteAction(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
