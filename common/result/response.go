package result

import "net/http"

type ResponseSuccess struct {
	Code uint32      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type ResponseFail struct {
	Code uint32 `json:"code"`
	Msg  string `json:"msg"`
}

func Success(data interface{}) *ResponseSuccess {
	return &ResponseSuccess{http.StatusOK, "OK", data}
}

func Fail(msg string) *ResponseFail {
	return &ResponseFail{
		Code: http.StatusBadRequest,
		Msg:  msg,
	}
}
