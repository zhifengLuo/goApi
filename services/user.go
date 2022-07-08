package services

import (
	"goapi/library"
	"goapi/models"
)

type User struct {
	Model *models.User
}

func (u *User) GetList(username, mobile string, pagination *library.Pagination) interface{} {
	return u.Model.GetList(username, mobile, pagination)
}

func (u *User) GetDetail(id int, username string) interface{} {
	if username != "" {
		return u.Model.GetByUsername(username)
	}
	return u.Model.GetById(id)
}

func (u *User) Register(user *models.UserReg) interface{} {
	return u.Model.Add(user)
}
