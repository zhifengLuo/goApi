package model

import (
	"time"
)

type UserToken struct {
	ID            int `gorm:"primarykey" json:"id"`
	Token         string
	TokenableType string
	TokenableId   int
	CreatedAt     LocalTime
	UpdatedAt     LocalTime
	ExpiredAt     LocalTime
	LastUsedAt    LocalTime
}

// TableName 指定表名
func (UserToken) TableName() string {
	return "user_token"
}

func (u *UserToken) GetTokenableId(token string) int {
	var row *UserToken
	db.Where("token = ?", token).First(&row)
	return row.TokenableId
}

func (u *UserToken) Get(token string) (row *UserToken) {
	db.Where("token = ?", token).First(&row)
	return
}

func (u *UserToken) Add(userToken *UserToken) bool {
	userToken.CreatedAt = LocalTime{Time: time.Now()}
	userToken.UpdatedAt = LocalTime{Time: time.Now()}
	userToken.LastUsedAt = LocalTime{Time: time.Now()}
	userToken.ExpiredAt = LocalTime{Time: time.Now().AddDate(0, 0, 10)}
	res := db.Create(&userToken)
	if res.RowsAffected == 0 {
		return false
	}
	return true
}

func (u *UserToken) Update(token string) bool {
	res := db.Model(&UserToken{}).Where("token = ?", token).Update("last_used_at", LocalTime{Time: time.Now()})
	if res.RowsAffected == 0 {
		return false
	}
	return true
}

func (u *UserToken) Remove(tokenableType string, tokenableId int) int {
	var rows []UserToken
	res := db.Where("tokenable_type = ? and tokenable_id = ?", tokenableType, tokenableId).Delete(&rows)
	return int(res.RowsAffected)
}
