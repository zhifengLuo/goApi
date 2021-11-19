package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goapi/controllers"
	"goapi/middleware"
	"net/http"
	"strings"
)

func NewRouter() *gin.Engine {
	// Default 使用 Logger 和 Recovery 中间件
	r := gin.Default()

	// 自定义中间件
	r.Use(cors())
	// router.Use(middleware.Auth())

	r.NoRoute(func(c *gin.Context) {
		c.JSON(400, gin.H{"code": 400, "error": "Bad Request"})
	})

	// 例子
	demo := new(controllers.DemoController)
	r.GET("/health", demo.Status)
	r.GET("/test", controllers.Test)
	r.GET("/json", controllers.DemoCtl.SendJson)

	// basic auth
	authorized := r.Group("/user")
	authorized.Use(middleware.AuthBasic())
	{
		authorized.GET("/info", controllers.UserCtl.Info)
	}

	// casbin auth
	auth := r.Group("/admin")
	auth.Use(middleware.NewAuthorizer())
	{
		auth.GET("/index", controllers.UserCtl.Index)
	}

	return r
}

// 跨域
func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method               // 请求方法
		origin := c.Request.Header.Get("Origin") // 请求头部
		var headerKeys []string                  // 声明请求头keys
		for k, _ := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", "*")                                       // 这是允许访问所有域
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE") // 服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			// 允许跨域设置
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar") // 跨域关键设置 让浏览器可以解析
			c.Header("Access-Control-Max-Age", "172800")                                                                                                                                                           // 缓存请求信息 单位为秒
			c.Header("Access-Control-Allow-Credentials", "false")                                                                                                                                                  //  跨域请求是否需要带cookie信息 默认设置为true
			c.Set("content-type", "application/json")                                                                                                                                                              // 设置返回格式是json
		}

		// 放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}

		// 处理请求
		c.Next()
	}
}
