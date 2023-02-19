package like

import (
	"net/http"

	"awesomeProject/dou-yin/like/cmd/api/internal/logic/like"
	"awesomeProject/dou-yin/like/cmd/api/internal/svc"
	"awesomeProject/dou-yin/like/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func LikeListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LikeListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := like.NewLikeListLogic(r.Context(), svcCtx)
		resp, err := l.LikeList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
