package router

import (
	"github.com/gin-gonic/gin"
	"goapi/controllers/api"
)

func groupApi(g *gin.RouterGroup) {
	g.GET("/user/info/:username", api.UserInfo)
	g.POST("/user/register", api.UserRegister)
}
