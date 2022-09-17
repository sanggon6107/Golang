package main

// mux : multiplexer. mux returns a http response corresponding to the input URL.
// When you have many handlers, you can use "mux"
// You can add handler to the mux object, and then when the server calls "http.ListenAndServe" fucntion,
// put the mux as the second arguemnt. mux functions as kind of "router"

import (
	"fmt"
	"net/http"
)

// how to parse query parameters
func IndexPathHandler(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query() // query parameters can be parsed through the object "r" since http.Request objects have the information regarding http requests.
	name := values.Get("name")
	age := values.Get("age")
	if name == "" {
		name = "Anonymous"
	}
	if age == "" {
		age = "?"
	}
	fmt.Fprintf(w, "Hi %s, who is %s years old! \n", name, age)
}

type GooHandler struct {
	goo string
	bar string
}

func (g *GooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, g.goo)
	fmt.Fprintln(w, g.bar)
}

func main() {

	// declare a mux object.
	mux := http.NewServeMux()

	mux.HandleFunc("/", IndexPathHandler)
	mux.Handle("/Goo", &GooHandler{"Kim", "dd"})

	http.ListenAndServe(":3000", mux)
}
