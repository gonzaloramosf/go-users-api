package models

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model

	Title       string `gorm:"type:varchar(200);not null;unique_index" json:"title"`
	Description string `json:"description"`
	Status      bool   `gorm:"default:false" json:"status"`
	UserID      uint   `gorm:"not null;unique_index" json:"userId"`
}
