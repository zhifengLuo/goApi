package controllers

import (
	"github.com/gin-gonic/gin"
)

type UserController struct{}

var UserCtl UserController

func (t UserController) Info(c *gin.Context) {

}

func (t UserController) Index(c *gin.Context) {
	success(c, "admin/index")
}
