package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Name    string `gorm:"not null"`
	Surname string `gorm:"not null"`
	Email   string `gorm:"not null;unique_index"`
	Tasks   []Task
}
