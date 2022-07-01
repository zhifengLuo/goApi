package api

import (
	"github.com/gin-gonic/gin"
	"goapi/controllers"
)

func GetUser(c *gin.Context) {
	id := c.Param("id")
	controllers.Success(c, "api "+id+" info")
}

func TestDone() {
	//n := models.NewEnforce()
}
