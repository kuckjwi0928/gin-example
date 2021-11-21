package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const DbName = "example.db"

var DB *gorm.DB

func Init() {
	db, err := gorm.Open(sqlite.Open(DbName), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("Failed to connect database")
	}
	_ = db.AutoMigrate(&Board{})
	DB = db
}
