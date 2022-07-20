package library

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"goapi/config"
	"strconv"
)

func NewRedis() *redis.Client {
	dbIndex, _ := strconv.Atoi(config.Get("REDIS_INDEX").(string))
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Get("REDIS_HOST"), config.Get("REDIS_PORT")),
		Password: fmt.Sprint(config.Get("REDIS_PWD")),
		DB:       dbIndex, // use default DB
	})
	return rdb
}
