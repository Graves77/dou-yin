package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"net/http"
)

// 限流频率为每秒5次请求（分别代表每秒产生的令牌数量和令牌桶大小）
var rl = ratelimit.NewBucketWithRate(5, 5)

func RateLimiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		if rl.TakeAvailable(1) < 1 {
			c.String(http.StatusForbidden, "请勿频繁访问")
			c.Abort()
			return
		}
		c.Next()
	}
}
