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

func UserLogin(c *gin.Context) {
	param := make(map[string]string)
	err := c.ShouldBindJSON(&param)

	if err != nil {
		controllers.Failure(c, 100, err.Error())
		return
	}
	if param["username"] == "" || param["password"] == "" {
		controllers.Failure(c, 100, "参数错误")
		return
	}
	service := services.User{}
	data := service.LoginByUsername(param["username"], param["password"])
	if data == 0 {
		controllers.Failure(c, 100, "账号不存在")
		return
	} else if data == 1 {
		controllers.Failure(c, 100, "密码错误")
		return
	}
	controllers.Success(c, data)
}
