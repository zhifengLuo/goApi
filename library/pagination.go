package library

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

type Pagination struct {
	Page      int         `form:"page" json:"page,omitempty;query:page"`
	PageSize  int         `form:"page_size" json:"page_size,omitempty;query:page_size"`
	Total     int64       `json:"total"`
	TotalPage int         `json:"total_page"`
	List      interface{} `json:"list"`
}

func NewPagination(c *gin.Context) *Pagination {
	var p = &Pagination{}
	page, _ := strconv.Atoi(c.Query("page"))
	if page == 0 {
		page = 1
	}
	p.Page = page
	pageSize, _ := strconv.Atoi(c.Query("page_size"))
	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}
	p.PageSize = pageSize
	return p
}
