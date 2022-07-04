package models

type User struct {
	Model
	Username string
	Nickname string
	Password string
	Mobile   string
	Sex      uint8 `json:"sex"`
	status   uint8
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

func (u *User) GetList(username, mobile string, page, pageSize int) interface{} {
	var total int64
	var rows *[]User
	data := make(map[string]interface{})
	query := db.Select("id,username,nickname,sex,status,mobile,created_at")
	if username != "" {
		query.Where("username like ?", username+"%")
	}
	if mobile != "" {
		query.Where("mobile like ?", mobile+"%")
	}
	query.Model(&rows).Count(&total)
	if page > 0 && pageSize > 0 {
		query.Offset((page - 1) * pageSize).Limit(pageSize)
	}
	query.Find(&rows)
	data["total"] = total
	data["list"] = rows
	return data
}
