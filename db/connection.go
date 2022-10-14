package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DSN = "host=localhost user=jack password=root dbname=gorm port=5432"
var DB *gorm.DB

func ConnectionDB() {
	var err error
	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("DB connected")
	}
}
