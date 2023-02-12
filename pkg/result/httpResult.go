package result

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

// HttpResult response
func HttpResult(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {
	if err == nil {
		//return success
		res := Success(resp)
		httpx.WriteJson(w, http.StatusOK, res)
	} else {
		logx.WithContext(r.Context())
		res := Fail(err.Error())
		httpx.WriteJson(w, http.StatusOK, res)
	}
}
