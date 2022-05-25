package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type DemoController struct{}

var DemoCtl = new(DemoController)

func TestOk(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

func (t DemoController) SendJson(c *gin.Context) {
	data := make(map[string]interface{})
	data["msg"] = "articles"
	data["list"] = []int{1, 3, 5, 7}
	data["total"] = 100
	success(c, data)
}
