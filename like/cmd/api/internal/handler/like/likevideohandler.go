package like

import (
	"net/http"

	"awesomeProject/dou-yin/like/cmd/api/internal/logic/like"
	"awesomeProject/dou-yin/like/cmd/api/internal/svc"
	"awesomeProject/dou-yin/like/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func LikeVideoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LikeReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := like.NewLikeVideoLogic(r.Context(), svcCtx)
		resp, err := l.LikeVideo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
