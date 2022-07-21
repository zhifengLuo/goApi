package api

import (
	"github.com/gin-gonic/gin"
	"goapi/controller"
	"goapi/model"
)

func Areas(c *gin.Context) {
	data := model.GetAllProvince()
	controller.Success(c, data)
}
