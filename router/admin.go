package router

import (
	"github.com/gin-gonic/gin"
	"goapi/controllers/admin"
)

func groupAdmin(g *gin.RouterGroup) {
	g.POST("/user/list", admin.UserList)
	g.POST("/user/detail/:id", admin.UserDetail)
}
