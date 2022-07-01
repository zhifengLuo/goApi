package router

import (
	"github.com/gin-gonic/gin"
	"goapi/controllers/api"
)

func groupApi(g *gin.RouterGroup) {
	g.POST("/user/:id", api.GetUser)
}
