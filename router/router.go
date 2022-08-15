package router

import (
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	// Default 使用 Logger 和 Recovery 中间件
	r := gin.Default()

	// 自定义中间件
	r.Use(cors())

	r.NoRoute(func(c *gin.Context) {
		c.JSON(400, gin.H{"code": 400, "error": "Bad Request"})
	})

	// 分组
	groupApi(r.Group("/api"))
	groupAdmin(r.Group("/admin"))

	return r
}

// 跨域
func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST ,GET, PUT, DELETE, UPDATE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, AccessToken, X-CSRF-Token, Authorization, Token, Range")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type, Content-Range")
		//c.Header("Access-Control-Max-Age", "172800")
		c.Header("Access-Control-Allow-Credentials", "true")
		//c.Header("Content-Range", "invoices: 0-24/120")

		method := c.Request.Method // 放行所有OPTIONS方法
		if method == "OPTIONS" {
			//c.JSON(http.StatusOK, "Options Request!")
			c.AbortWithStatus(200)
		}

		// 处理请求
		c.Next()
	}
}
