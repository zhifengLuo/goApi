package admin

import (
	"github.com/gin-gonic/gin"
	"goapi/controller"
	"goapi/library"
	"goapi/service"
	"strconv"
)

func UserList(c *gin.Context) {
	username := c.PostForm("username")
	mobile := c.PostForm("mobile")
	service := service.UserService{}
	pagination := library.NewPagination(c)
	data := service.GetList(username, mobile, pagination)
	controller.Success(c, data)
}

func UserDetail(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	username := c.Query("username")
	service := service.UserService{}
	data := service.GetDetail(id, username)
	controller.Success(c, data)
}

func UserSetPassword(c *gin.Context) {

}
