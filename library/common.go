package library

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func ToDateTime(t time.Time) string {
	if t.Unix() == 0 {
		return ""
	}
	return fmt.Sprintf("%d-%d-%d %02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
}

// 返回随机字符串
func RandomStr(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// 数组切片去重
func ArrayUnique(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

// 将字符串分割为int数组
func SplitToInt(str, sep string) ([]int, error) {
	s := strings.Split(str, sep)
	r := make([]int, 0, len(s))
	for _, v := range s {
		t, err := strconv.Atoi(strings.Trim(v, " "))
		if err != nil {
			continue
		}
		r = append(r, t)
	}

	return r, nil
}
