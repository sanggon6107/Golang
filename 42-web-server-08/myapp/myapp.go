package myapp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	CreatedAt time.Time `json:"created_at"`
}

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

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}

	user.ID = 2
	user.CreatedAt = time.Now()
	w.WriteHeader(http.StatusCreated)
	data, _ := json.Marshal(user) // to json format
	fmt.Fprint(w, string(data))
}

func NewHandler() http.Handler {
	mux := mux.NewRouter()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/users", usersHandler).Methods("GET")
	mux.HandleFunc("/users", createUserHandler).Methods("POST")
	mux.HandleFunc("/users/{id:[0-9]+}", getUsersHandler)

	return mux
}
