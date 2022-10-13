package main

import (
	"github.com/gorilla/mux"
	"go-users-api/router"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", router.HomeHandler)

	http.ListenAndServe(":8080", r)
}
