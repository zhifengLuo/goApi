package controllers

import (
	"github.com/gin-gonic/gin"
	"goapi/models"
)

func GetUser(c *gin.Context) {
	id := c.Param("id")
	success(c, "user "+id+" info")
}

func Areas(c *gin.Context) {
	data := models.GetAllProvince()
	success(c, data)
}
