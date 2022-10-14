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

	// / route
	r.HandleFunc("/", router.Home)

	// users routes
	r.HandleFunc("/users", router.GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", router.GetUser).Methods("GET")
	r.HandleFunc("/users", router.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", router.DeleteUser).Methods("DELETE")

	// tasks routes
	r.HandleFunc("/tasks", router.GetTasks).Methods("GET")
	r.HandleFunc("/tasks/{id}", router.GetTask).Methods("GET")
	r.HandleFunc("/tasks", router.CreateTask).Methods("POST")
	r.HandleFunc("/tasks/{id}", router.DeleteTask).Methods("DELETE")

	http.ListenAndServe(":8080", r)
}
