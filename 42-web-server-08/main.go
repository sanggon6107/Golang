package main

// RESTful API : Representational State Transfer.
// It is a sort of protocol(interface or standard) to make a secured information exchange.

// GET /users : get a list of users
// GET /users/:id : find a particular user in users
// POST /users : create a new user
// PUT /users : update an existing user
// DELETE : delete an existing user

import (
	"net/http"

	"github.com/sanggon6107/Golang/tree/master/42-web-server-08/myapp"
)

func main() {

	http.ListenAndServe(":3000", myapp.NewHandler())
}
