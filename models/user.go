package models

type User struct {
	Model
	Username string
	Password string
}

// 指定表名
/*func (u User) TableName() string {
	return "user"
}*/
