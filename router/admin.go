package router

import (
	"github.com/gin-gonic/gin"
	"goapi/controllers/admin"
)

func groupAdmin(g *gin.RouterGroup) {
	g.GET("/user/list", admin.ListUsers)
}
