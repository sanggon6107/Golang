package main

import "net/http"

func main() {

	http.ListenAndServe("/", myapp.NewHandler())
}