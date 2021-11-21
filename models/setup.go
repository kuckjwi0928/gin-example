package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	DBName = "example.db"
	TestDB = ":memory:"
)

var database *gorm.DB

func Init() {
	db, err := gorm.Open(sqlite.Open(DBName), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("Failed to connect database")
	}
	_ = db.AutoMigrate(&Board{})
	database = db
}

func TestInit() {
	db, _ := gorm.Open(sqlite.Open(TestDB), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	_ = db.AutoMigrate(&Board{})
	database = db
}

func GetDB() *gorm.DB {
	return database
}
