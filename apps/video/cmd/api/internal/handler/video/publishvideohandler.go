package video

import (
	"awesomeProject/dou-yin/apps/video/cmd/api/internal/logic/video"
	"awesomeProject/dou-yin/apps/video/cmd/api/internal/svc"
	"awesomeProject/dou-yin/apps/video/cmd/api/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func PublishVideoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PublishVideoActionReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := video.NewPublishVideoLogic(r.Context(), svcCtx)
		resp, err := l.PublishVideo(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
