package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type User struct {
	Name  string    `json:"name"` // matches names between json <-> Golang. annotation.
	Age   string    `json:"age"`
	Email string    `json:"email"`
	Time  time.Time `json:"time"`
}

type FooHandler struct{}

func (f *FooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	// NewDecoder receives "io.Reader" objects as a parameter, and r.Body is an implementation of io.Reader since it has a method called "Read".
	// NewDecoder.Decode function returns a decoded object of json.
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	user.Time = time.Now()

	// json.Marshal encodes the argument as []byte(json), and then returns it.
	// This means that if I want to use this as string, I would need to cast it.
	data, _ := json.Marshal(user)
	w.Header().Add("content-type", "application/json") // To let the request header know "data" is a json format.
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, string(data))

}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Index")
}

// returns http.Handler(mux)
func Handler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", IndexHandler)
	mux.Handle("/Foo", &FooHandler{})

	return mux
}

func main() {

	err := http.ListenAndServe(":3000", Handler())
	if err != nil {
		fmt.Println("err")
	}
}
