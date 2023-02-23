package middleware

import (
	"awesomeProject/dou-yin/apps/video/cmd/rpc/types/video"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UploadFileMiddleware struct {
}

func NewUploadFileMiddleware() *UploadFileMiddleware {
	return &UploadFileMiddleware{}
}

func (m *UploadFileMiddleware) Handle(c *gin.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		file, header, _ := c.Request.FormFile("data")
		video.SaveFile(file, header)
		token := c.PostForm("token")
		title := c.PostForm("title")
		fmt.Println(token, title)
		//service.UploadService(param1, param2, filename)
		c.String(http.StatusOK, "File uploaded successfully")

	}
}
