package myapp

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func TestNewHandler(t *testing.T) {
	assert := assert.New(t)
	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	resp, err := http.Get(ts.URL) // get response by "Get"
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	data, _ := ioutil.ReadAll(resp.Body)
	assert.Equal("Hello world!", string(data))
}

func TestUsers(t *testing.T) {
	assert := assert.New(t)
	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/users")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	data, _ := ioutil.ReadAll(resp.Body) // response
	assert.Contains(string(data), "Get UserInfo")
}

// test for "Get"
func TestgetUsersHandler(t *testing.T) {
	assert := assert.New(t)
	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	userId := strconv.Itoa(rand.Intn(99))
	resp, err := http.Get(ts.URL + "/users/" + userId)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	data, _ := ioutil.ReadAll(resp.Body) // response
	assert.Contains(string(data), "User Id:"+userId)
}

// test for "POST"
func TestCreateUserHandler(t *testing.T) {
	assert := assert.New(t)
	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	resp, err := http.Post(ts.URL+"/users", "application/json",
		strings.NewReader(`{"fist_name":"salcan", "last_name":"Han"}`))
	assert.NoError(err)
	assert.Equal(http.StatusCreated, resp.StatusCode)

	user := new(User)
	err = json.NewDecoder(resp.Body).Decode(user)
	assert.NoError(err)
	assert.NotEqual(0, user.ID)

	id := user.ID
	// to make sure that the id of user created by "POST" equals the id from "GET"
	resp, err = http.Get(ts.URL + "/users/" + strconv.Itoa(id))
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	user2 := new(User)
	err = json.NewDecoder(resp.Body).Decode(user2) // read the data inside the body of the response.
	assert.NoError(err)
	assert.Equal(user.ID, user2.ID)
	assert.Equal(user.FirstName, user2.FirstName)
}
