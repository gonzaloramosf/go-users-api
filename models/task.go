package models

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model

	Title       string `gorm:"type:varchar(200);not null;unique_index"`
	Description string
	Status      bool `gorm:"default:false"`
	UserID      uint `gorm:"not null;unique_index"`
}
