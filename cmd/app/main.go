package main

import (
	"github.com/gorilla/mux"
	"go-users-api/db"
	"go-users-api/models"
	"go-users-api/router"
	"net/http"
)

func main() {
	db.ConnectionDB()

	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})

	r := mux.NewRouter()
	r.HandleFunc("/", router.HomeHandler)
	http.ListenAndServe(":8080", r)
}
