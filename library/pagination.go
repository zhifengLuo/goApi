package library

import (
	"math"
)

type Pagination struct {
	page      int         `form:"page" json:"page,omitempty;query:page"`
	pageSize  int         `form:"page_size" json:"limit,omitempty;query:page_size"`
	total     int64       `json:"total"`
	totalPage int         `json:"total_page"`
	list      interface{} `json:"list"`
}

func (p *Pagination) GetOffset() int {
	return (p.page - 1) * p.pageSize
}

func (p *Pagination) TotalPage() int {
	return int(math.Ceil(float64(p.total) / float64(p.pageSize)))
}

func (p *Pagination) Page() int {
	if p.page == 0 {
		p.page = 1
	}
	return p.page
}

func (p *Pagination) PageSize() int {
	if p.pageSize == 0 {
		p.pageSize = 10
	}
	return p.pageSize
}

func (p *Pagination) SetPage(page int) {
	p.page = page
}

func (p *Pagination) SetPageSize(pageSize int) {
	p.pageSize = pageSize
}
