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
	var user *models.User
	if username != "" {
		user = u.Model.GetByUsername(username)
	} else {
		user = u.Model.GetById(id)
	}
	user.Mobile = user.GetMobile()
	return user
}

func (u *User) Register(user *models.UserReg) interface{} {
	return u.Model.Add(user)
}

func (u *User) LoginByUsername(username, password string) interface{} {
	var user *models.User
	user = u.Model.GetByUsername(username)
	if user.ID == 0 {
		return 0
	}
	res := u.Model.CheckPassword(password, user.Password)
	if res {
		return user
	}
	return 1
}

func (u *User) LoginByMoible(mobile, code string) interface{} {
	return true
}
