package admin

import (
	"github.com/gin-gonic/gin"
	"goapi/controllers"
)

func ListUsers(c *gin.Context) {
	data := []string{"abc", "123"}
	controllers.Success(c, data)
}
