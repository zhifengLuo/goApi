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
	service := service.UserService{}
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
	service := service.UserService{}
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
	sUser := service.UserService{}
	data, msg := sUser.LoginByUsername(param["username"], param["password"])
	if data == nil {
		controller.Failure(c, 100, msg)
		return
	}
	sToken := service.UserTokenSevice{}
	token := sToken.CreateToken(service.TypeUser, data.ID)
	res := make(map[string]interface{})
	res["token"] = token
	res["user"] = data
	controller.Success(c, res)
}

func TestSet(c *gin.Context) {
	vars := 456
	controller.Success(c, vars)
}

func TestGet(c *gin.Context) {
	e := model.NewEnforce()
	sub := c.Query("sub")
	obj := c.Query("obj")
	act := c.Query("act")

	res, _ := e.Enforce(sub, obj, act)
	controller.Success(c, res)
}
