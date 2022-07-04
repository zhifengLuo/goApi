package services

import "goapi/models"

type User struct {
	Service *models.User
}

func (u *User) GetList(username, mobile string, page, pageSize int) interface{} {
	return u.Service.GetList(username, mobile, page, pageSize)
}
