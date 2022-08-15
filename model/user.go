package model

import (
	"fmt"
	"goapi/library"
	"golang.org/x/crypto/bcrypt"
	"time"
)

var SexTxt = []string{1: "男", 2: "女", 3: "其他"}

type User struct {
	Model
	Username string `gorm:"unique"`
	Nickname string
	Password string `json:"-"`
	Mobile   string
	Sex      uint8 `json:"sex"`
	Status   uint8
}

type UserReg struct {
	Username string `binding:"required"`
	Nickname string
	Mobile   string `binding:"required"`
	Password string `binding:"required"`
	Sex      uint8
}

// TableName 指定表名
func (User) TableName() string {
	return "user"
}

// Check 验证用户
func (u User) Check() {

}

func (u *User) GetStatusTxt() string {
	status := []string{"禁用", "锁定", "启用"}
	return status[u.Status]
}

func (u *User) GetMobile() string {
	s := []rune(u.Mobile)
	total := len(s)
	if total == 0 {
		return ""
	}
	return string(s[0:3]) + "****" + string(s[7:11])
}

func (u *User) SetPassword(password string) {
	newPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	u.Password = string(newPassword)
}

func (u *User) CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (u *User) Add(reg *UserReg) interface{} {
	var user User
	user.Username = reg.Username
	user.Nickname = reg.Nickname
	user.Mobile = reg.Mobile
	user.SetPassword(reg.Password)
	user.Sex = reg.Sex
	user.Status = 1
	user.CreatedAt = LocalTime{Time: time.Now()}
	user.UpdatedAt = LocalTime{Time: time.Now()}
	result := db.Create(&user)
	if result.RowsAffected == 0 {
		fmt.Println(result.Error)
	}
	return user.ID
}

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
	pagination.Data = rows
	return pagination
}

func (u *User) GetById(id int) *User {
	var row User
	db.First(&row, id)
	return &row
}

func (u *User) GetByUsername(username string) *User {
	var row User
	db.Where("username = ?", username).First(&row)
	return &row
}

func (u *User) GetByMobile(mobile string) *User {
	var row User
	db.Where("mobile = ?", mobile).First(&row)
	return &row
}
