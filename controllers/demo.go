package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type DemoController struct{}

var DemoCtl DemoController

func (t DemoController) Status(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

func (t DemoController) SendJson(c *gin.Context) {
	data := make(map[string]interface{})
	data["msg"] = "articles"
	data["list"] = []int{1, 3, 5, 7}
	data["total"] = 100
	success(c, data)
}

func Test(c *gin.Context) {
	c.JSON(200, gin.H{"code": 0, "message": "demo"})
}
