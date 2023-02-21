package middleware

import "net/http"

type UploadFileMiddleware struct {
}

func NewUploadFileMiddleware() *UploadFileMiddleware {
	return &UploadFileMiddleware{}
}

func (m *UploadFileMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		next(w, r)
	}
}
