package admin

import (
	"github.com/gin-gonic/gin"
	"goapi/controllers"
	"goapi/services"
	"strconv"
)

func UserList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultPostForm("page", "0"))
	pageSize, _ := strconv.Atoi(c.PostForm("page_size"))
	username := c.PostForm("username")
	mobile := c.PostForm("mobile")
	service := services.User{}
	data := service.GetList(username, mobile, page, pageSize)
	controllers.Success(c, data)
}

func UserDetail(c *gin.Context) {

}

func UserSetPassword(c *gin.Context) {

}
