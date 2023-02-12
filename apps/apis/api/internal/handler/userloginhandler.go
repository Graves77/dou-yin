package handler

import (
	"awesomeProject/dou-yin/apps/apis/api/internal/logic"
	"awesomeProject/dou-yin/apps/apis/api/internal/svc"
	"awesomeProject/dou-yin/apps/apis/api/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func userLoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUserLoginLogic(r.Context(), svcCtx)
		resp, err := l.UserLogin(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
