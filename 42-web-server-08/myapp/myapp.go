package myapp

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func usersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Get UserInfo by /users/{id}")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world!")
}

func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, "User Id:", vars["id"])
}

func NewHandler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/users", usersHandler)
	mux.HandleFunc("/users/{id:[0-9]+}", getUsersHandler)

	return mux
}
