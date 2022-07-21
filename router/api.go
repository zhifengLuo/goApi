package router

import (
	"github.com/gin-gonic/gin"
	"goapi/controller/api"
	"goapi/middleware"
)

func groupApi(g *gin.RouterGroup) {
	g.GET("/user/info/:username", api.UserInfo)
	g.GET("/test/get", api.TestGet)
	g.POST("/user/register", api.UserRegister)
	g.POST("/user/login", api.UserLogin)

	g.Use(middleware.AuthUser())
	{
		g.GET("/test/set", api.TestSet)
	}
}
