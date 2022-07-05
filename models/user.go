package models

import "goapi/library"

type User struct {
	Model
	Username string
	Nickname string
	password string
	Mobile   string
	Sex      uint8 `json:"sex"`
	Status   uint8
}

// TableName 指定表名
func (u User) TableName() string {
	return "user"
}

// Check 验证用户
func (u User) Check() {

}

// 注册
// 登入
// 忘记密码
// 修改密码
// 更新信息

func (u *User) GetList(username, mobile string, pagination *library.Pagination) *library.Pagination {
	var rows *[]User
	query := db.Select("id,username,nickname,sex,status,mobile,created_at,updated_at")
	if username != "" {
		query.Where("username like ?", username+"%")
	}
	if mobile != "" {
		query.Where("mobile like ?", mobile+"%")
	}
	query.Order("id desc")
	query.Scopes(paginate(rows, pagination)).Find(&rows)
	pagination.List = rows
	return pagination
}
