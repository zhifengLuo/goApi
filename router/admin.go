package router

import (
	"github.com/gin-gonic/gin"
	"goapi/controller/admin"
)

func groupAdmin(g *gin.RouterGroup) {
	g.GET("/user/list", admin.UserList)
	g.GET("/user/detail/:id", admin.UserDetail)
}
