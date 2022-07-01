package api

import (
	"github.com/gin-gonic/gin"
	"goapi/controllers"
	"goapi/models"
)

func Areas(c *gin.Context) {
	data := models.GetAllProvince()
	controllers.Success(c, data)
}
