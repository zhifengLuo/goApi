package models

type User struct {
	Model
	Username string
	Password string
	Mobile   string
	Sex      uint8
	Nickname string
	Avatar   string
}

// 指定表名
func (u User) TableName() string {
	return "users"
}

// 验证用户
func (u User) Check() {

}

// 注册
// 登入
// 忘记密码
// 修改密码
// 更新信息
