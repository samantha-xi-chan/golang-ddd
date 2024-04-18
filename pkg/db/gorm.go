package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 数据库配置
func NewDB() (*gorm.DB, error) {
	dsn := "root:gznPzkTJ8xEgGZO6@tcp(127.0.0.1:3306)/biz?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
