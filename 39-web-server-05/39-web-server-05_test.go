package main

// The name of the source code for testing has to be ***_test.go
// the name of the function has to be Test***, and it has to have a *testing.T parameter called t.

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestIndexHandler(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil) // send a http request for testing.

	mux := Handler()
	mux.ServeHTTP(res, req) // This "router" would find the handler crresponding to "/" path, and send http request (req).

	assert.Equal(http.StatusOK, res.Code)
	data, _ := io.ReadAll(res.Body)           // read the body of the result.
	assert.Equal("Hello world", string(data)) // fail : "Hello world", not "Index" as written in "39-webs-server-05.go"
}

func TestFooHandler(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/Foo",
		strings.NewReader(`{"name":"Han", "age":"20", "email":"han@go.com"}`)) // send a http request for testing.

	mux := Handler()
	mux.ServeHTTP(res, req) // This "router" would find the handler crresponding to "/" path, and send http request (req).

	assert.Equal(http.StatusCreated, res.Code)

	user := new(User)
	err := json.NewDecoder(res.Body).Decode(user)
	assert.Nil(err) // err should be nil
	assert.Equal("Han", user.Name)
	assert.Equal("20", user.Age)
	assert.Equal("han@go.com", user.Email)
}
