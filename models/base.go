package models

import (
	"fmt"
	"goapi/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

type Model struct {
	gorm.Model
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

func Get(result interface{}, where map[string]interface{}) interface{} {
	db.Where(where).First(&result)
	return result
}

func GetAll(result interface{}, where map[string]interface{}) interface{} {
	db.Where(where).Find(&result)
	return result
}
