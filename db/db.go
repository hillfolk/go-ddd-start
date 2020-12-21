package db

import "gorm.io/gorm"

var db *gorm.DB

func Init() {
}

func GetDB() *gorm.DB {
	return db
}
