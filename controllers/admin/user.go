package admin

import (
	"github.com/gin-gonic/gin"
	"goapi/controllers"
	"goapi/library"
	"goapi/services"
	"strconv"
)

func UserList(c *gin.Context) {
	username := c.PostForm("username")
	mobile := c.PostForm("mobile")
	service := services.User{}
	pagination := library.NewPagination(c)
	data := service.GetList(username, mobile, pagination)
	controllers.Success(c, data)
}

func UserDetail(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	username := c.Query("username")
	service := services.User{}
	data := service.GetDetail(id, username)
	controllers.Success(c, data)
}

func UserSetPassword(c *gin.Context) {

}
