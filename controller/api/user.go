package api

import (
	"github.com/gin-gonic/gin"
	"goapi/controller"
	"goapi/model"
	"goapi/service"
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
	service := service.User{}
	data := service.GetDetail(0, username)
	controller.Success(c, data)
}

func UserRegister(c *gin.Context) {
	var param *model.UserReg
	err := c.ShouldBindJSON(&param)

	if err != nil {
		controller.Failure(c, 100, err.Error())
		return
	}
	service := service.User{}
	data := service.Register(param)
	controller.Success(c, data)
}

func UserLogin(c *gin.Context) {
	param := make(map[string]string)
	err := c.ShouldBindJSON(&param)

	if err != nil {
		controller.Failure(c, 100, err.Error())
		return
	}
	if param["username"] == "" || param["password"] == "" {
		controller.Failure(c, 100, "参数错误")
		return
	}
	sUser := service.User{}
	data, msg := sUser.LoginByUsername(param["username"], param["password"])
	if data == nil {
		controller.Failure(c, 100, msg)
		return
	}
	sToken := service.UserToken{}
	sToken.CreateToken(service.TypeUser, data.ID)
	controller.Success(c, data)
}

func TestSet(c *gin.Context) {
	vars := 456
	controller.Success(c, vars)
}

func TestGet(c *gin.Context) {
	vars := 123
	controller.Success(c, vars)
}
