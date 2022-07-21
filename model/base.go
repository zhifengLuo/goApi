package model

import (
	"database/sql/driver"
	"fmt"
	"goapi/config"
	"goapi/library"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"math"
	"strings"
	"time"
)

var db *gorm.DB

type Model struct {
	ID        int `gorm:"primarykey" json:"id"`
	CreatedAt LocalTime
	UpdatedAt LocalTime
	deletedAt gorm.DeletedAt `gorm:"index"`
}

func init() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local", config.Get("DB_USER"), config.Get("DB_PWD"), config.Get("DB_HOST"), config.Get("DB_PORT"), config.Get("DB_NAME"), config.Get("DB_CHARSET"))
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			// TablePrefix: "tb_",
			SingularTable: true, // 使用单数表名
		},
		// Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		panic("failed to connect database")
	}
}

func paginate(value interface{}, pagination *library.Pagination) func(db *gorm.DB) *gorm.DB {
	var total int64
	db.Model(value).Count(&total)
	pagination.Total = total
	totalPages := int(math.Ceil(float64(total) / float64(pagination.PageSize)))
	pagination.TotalPage = totalPages
	offset := (pagination.Page - 1) * pagination.PageSize
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(offset).Limit(pagination.PageSize)
	}
}

type LocalTime struct {
	time.Time
}

func (t *LocalTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	var err error
	str := string(data)
	timeStr := strings.Trim(str, "\"")
	t1, err := time.Parse("2006-01-02 15:04:05", timeStr)
	*t = LocalTime{t1}
	return err
}
func (t LocalTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05"))
	return []byte(formatted), nil
}
func (t LocalTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}
func (t *LocalTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = LocalTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
