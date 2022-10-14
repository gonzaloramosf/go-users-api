package router

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go-users-api/db"
	"go-users-api/models"
	"net/http"
)

/* / routes */

func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("My first Go Api!"))
}

/* /users routes */

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	db.DB.Find(&users)

	for i := 0; i < len(users); i++ {
		db.DB.Model(&users[i]).Association("Tasks").Find(&users[i].Tasks)
	}

	json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	params := mux.Vars(r)
	db.DB.First(&user, params["id"])

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}

	db.DB.Model(&user).Association("Tasks").Find(&user.Tasks)

	json.NewEncoder(w).Encode(user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	json.NewDecoder(r.Body).Decode(&user)

	createdUser := db.DB.Create(&user)
	err := createdUser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	params := mux.Vars(r)

	db.DB.First(&user, params["id"])
	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User to delete not exist"))
		return
	}

	db.DB.Unscoped().Delete(user)
	w.WriteHeader(http.StatusOK)
}

/* /tasks routes */

func GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks := models.Task{}
	db.DB.Find(&tasks)
	json.NewEncoder(w).Encode(tasks)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	task := models.Task{}
	json.NewDecoder(r.Body).Decode(&task)
	createdTask := db.DB.Create(&task)

	err := createdTask.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(task)
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	params := mux.Vars(r)

	db.DB.First(&task, params["id"])
	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Task not found"))
		return
	}

	json.NewEncoder(w).Encode(task)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	task := models.Task{}
	params := mux.Vars(r)

	db.DB.First(&task, params["id"])
	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Task to delete not exist"))
		return
	}

	db.DB.Unscoped().Delete(task)
	w.WriteHeader(http.StatusOK)
}
