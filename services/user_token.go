package services

import (
	"context"
	"encoding/json"
	"goapi/library"
	"goapi/models"
	"time"
)

const (
	CacheKey     = "token:"
	TypeUser     = "user"
	TypeAdmin    = "admin"
	TypeEmployee = "employee"
)

type UserToken struct {
	Model *models.UserToken
}

func (u *UserToken) TokenCreate(tokenableType string, tokenableId int) string {
	tokenType := TypeUser
	var row models.UserToken
	switch tokenableType {
	case TypeAdmin:
		tokenType = TypeAdmin
	case TypeEmployee:
		tokenType = TypeEmployee
	}
	if tokenableId < 1 {
		return ""
	}
	row.Token = library.Md5(tokenType + string(tokenableId) + time.Now().String())
	row.TokenableType = tokenType
	row.TokenableId = tokenableId
	if u.Model.Add(&row) {
		// save cache
		redis := library.NewRedis()
		ctx := context.Background()
		value, _ := json.Marshal(row)
		redis.Set(ctx, CacheKey+row.Token, value, 9*24*60*time.Minute)
		defer redis.Close()
		return row.Token
	}
	return ""
}

func (u *UserToken) TokenInfo(token string) (row *models.UserToken) {
	redis := library.NewRedis()
	ctx := context.Background()
	value, _ := redis.Get(ctx, CacheKey+token).Result()
	if value == "" {
		// get from db
		row = u.Model.Get(token)
	} else {
		json.Unmarshal([]byte(value), &row)
	}
	defer redis.Close()
	return row
}

func (u *UserToken) GetUser(token string) {

}

func (u *UserToken) GetAdmin(token string) {

}
