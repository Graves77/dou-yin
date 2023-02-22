package video

import (
	"awesomeProject/dou-yin/apps/api/internal/logic/video"
	"awesomeProject/dou-yin/apps/api/internal/svc"
	"awesomeProject/dou-yin/apps/api/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func PublishVideoListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PublishVideoListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := video.NewPublishVideoListLogic(r.Context(), svcCtx)
		resp, err := l.PublishVideoList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
