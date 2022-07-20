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

func (u *User) LoginByUsername(username, password string) (user *models.User, msg string) {
	user = u.Model.GetByUsername(username)
	if user.ID == 0 {
		return nil, "账号不存在"
	}
	res := u.Model.CheckPassword(password, user.Password)
	if res {
		return user, ""
	}
	return nil, "密码错误"
}

func (u *User) LoginByMoible(mobile, code string) interface{} {
	return true
}
