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
	data, msg := service.LoginByUsername(param["username"], param["password"])
	if data == nil {
		controllers.Failure(c, 100, msg)
		return
	}
	sToken := services.UserToken{}
	sToken.CreateToken(services.TypeUser, data.ID)
	controllers.Success(c, data)
}

func TestSet(c *gin.Context) {
	vars := 456
	controllers.Success(c, vars)
}

func TestGet(c *gin.Context) {
	vars := 123
	controllers.Success(c, vars)
}
