package models

import "gorm.io/gorm"

type Board struct {
	gorm.Model
	Title   string `json:"title" binding:"required" gorm:"not null"`
	Content string `json:"content" binding:"required" gorm:"not null"`
}
