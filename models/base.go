package models

import (
	"fmt"
	"goapi/config"
	"goapi/library"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"math"
	"time"
)

var db *gorm.DB

type Model struct {
	ID        uint `gorm:"primarykey" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
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
