package api

import (
	"github.com/gin-gonic/gin"
	"goapi/controllers"
	"goapi/models"
	"goapi/services"
)

type UserReg struct {
	Username string
	Nickname string
	Mobile   string
	Password string
	Sex      uint8
}

func UserInfo(c *gin.Context) {
	username := c.Param("username")
	service := services.User{}
	data := service.GetDetail(0, username)
	controllers.Success(c, data)
}

func UserRegister(c *gin.Context) {
	var param *models.UserReg
	err := c.ShouldBindJSON(&param)

	if err != nil {
		controllers.Failure(c, 100, err.Error())
		return
	}
	service := services.User{}
	data := service.Register(param)
	controllers.Success(c, data)
}
