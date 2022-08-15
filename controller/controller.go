package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Success(c *gin.Context, data interface{}) {
	obj := make(map[string]interface{})
	obj["code"] = 0
	switch data.(type) {
	case string, int, float64, bool:
		obj["result"] = data
	default:
		result, _ := json.Marshal(data)
		json.Unmarshal(result, &obj)
	}
	c.JSON(http.StatusOK, obj)
}

func Failure(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"code":  code,
		"error": msg,
	})
}
