package router

import "net/http"

// HomeHandler / routes
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World Go!"))
}

// UserHandler /user routes
func UserHandler() {

}

// TaskHandler /task routes
func TaskHandler() {

}
