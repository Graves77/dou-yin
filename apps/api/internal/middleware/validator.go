package middleware

import (
	"awesomeProject/dou-yin/common/variable"
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
)

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Register struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginValidationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var login Login

		if err := c.ShouldBindJSON(&login); err != nil {
			variable.ZapLog.Info(err.Error())
			c.Abort()
			return
		}
		variable.ZapLog.Info(login.Username)
		if err := validation.ValidateStruct(&login,
			validation.Field(&login.Username, validation.Required),
			validation.Field(&login.Password, validation.Required),
		); err != nil {
			variable.ZapLog.Info(err.Error())
			c.Abort()
			return
		}
		c.Set("login", login)
		c.Next()
	}
}
