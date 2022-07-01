package middleware

import (
	"github.com/gin-gonic/gin"
)

func AuthBasic() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, password, status := c.Request.BasicAuth()
		errMsg := "The authentication is failed"
		if !status {
			c.AbortWithStatusJSON(401, gin.H{"code": 401, "error": errMsg})
			return
		}
		if !validateUser(username, password) {
			c.AbortWithStatusJSON(401, gin.H{"code": 401, "error": errMsg})
			return
		}
		c.Next()
	}
}

func validateUser(username string, password string) bool {
	if username == "foo" && password == "bar123" {
		return true
	}
	return false
}
