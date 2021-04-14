package controllers

import (
	"github.com/gin-gonic/gin"
)

type UserController struct{}

var UserCtl UserController

func (t UserController) Info(c *gin.Context) {}
