package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(dsn string) (*gorm.DB, error) {
	//conf := Get()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err == nil {
		DB = db
		return db, err
	}
	return nil, err
}
