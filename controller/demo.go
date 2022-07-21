package controller

import (
	"github.com/gin-gonic/gin"
	"goapi/config"
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
	Success(c, data)
}

func (t DemoController) TestConfig(c *gin.Context) {
	data := config.Get("Taxonomies")
	Success(c, data)
}
