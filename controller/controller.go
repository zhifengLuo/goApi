package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": data,
	})
}

func Failure(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"code":  code,
		"error": msg,
	})
}
