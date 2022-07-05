package services

import (
	"goapi/library"
	"goapi/models"
)

type User struct {
	Service *models.User
}

func (u *User) GetList(username, mobile string, pagination *library.Pagination) interface{} {
	return u.Service.GetList(username, mobile, pagination)
}
