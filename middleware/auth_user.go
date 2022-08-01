package middleware

import (
	"github.com/gin-gonic/gin"
	"goapi/service"
	"strings"
)

func AuthUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.Request.Header.Get("Authorization")
		tokenArr := strings.Split(tokenStr, "Bearer ")
		if len(tokenArr) != 2 {
			c.AbortWithStatusJSON(401, gin.H{"code": 401, "error": "access deny"})
		} else {
			token := tokenArr[1]
			s := service.UserTokenSevice{}
			if !s.CheckUser(token) {
				c.AbortWithStatusJSON(401, gin.H{"code": 401, "error": "access denied"})
			}
		}
	}
}
