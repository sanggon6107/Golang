package myapp

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/stretchr/testify/assert"

	// "math/rand"

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

// test for "GET"
func TestGetUsersHandler(t *testing.T) {
	assert := assert.New(t)
	ts := httptest.NewServer(NewHandler())
	defer ts.Close()
	// userId := strconv.Itoa(rand.Intn(99))
	resp, err := http.Get(ts.URL + "/users/99")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	data, _ := ioutil.ReadAll(resp.Body)
	assert.Contains(string(data), "No User Id:99")
}

// test for "POST"
func TestCreateUserHandler(t *testing.T) {
	assert := assert.New(t)
	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	// "POST"
	resp, err := http.Post(ts.URL+"/users", "application/json",
		strings.NewReader(`{"first_name":"Salcan", "last_name":"Han"}`))
	assert.NoError(err)
	assert.Equal(http.StatusCreated, resp.StatusCode)

	// parsing the response to the above code.
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
	assert.Equal(user.LastName, user2.LastName)
}

// test for "DELETE"
func TestDeleteUserHandler(t *testing.T) {
	assert := assert.New(t)
	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	// should be no user to delete
	req, _ := http.NewRequest("DELETE", ts.URL+"/users/1", nil)
	resp, err := http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	data, _ := ioutil.ReadAll(resp.Body)
	assert.Contains(string(data), "No user Id:1")

	// POST to make a user
	resp, err = http.Post(ts.URL+"/users", "application/json",
		strings.NewReader(`{"first_name":"Salcan", "last_name":"Han"}`))
	assert.NoError(err)
	assert.Equal(http.StatusCreated, resp.StatusCode)

	// parsing the response to the above code.
	user := new(User)
	err = json.NewDecoder(resp.Body).Decode(user)
	assert.NoError(err)
	assert.NotEqual(0, user.ID)

	idStr := strconv.Itoa(user.ID)
	req, _ = http.NewRequest("DELETE", ts.URL+"/users/"+idStr, nil)
	resp, err = http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	data, _ = ioutil.ReadAll(resp.Body)
	assert.Contains(string(data), "User deleted Id:"+idStr)
}

func TestUpdateUserHandler(t *testing.T) {
	assert := assert.New(t)
	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	resp, err := http.Post(ts.URL+"/users", "application/json",
		strings.NewReader(`{"first_name":"Salcan", "last_name":"Han"}`))
	assert.NoError(err)
	assert.Equal(http.StatusCreated, resp.StatusCode)
	userBeforeUpdate := new(User)
	err = json.NewDecoder(resp.Body).Decode(userBeforeUpdate)
	assert.NoError(err)

	jsonUpdate := fmt.Sprintf(`{"ID":%d, "first_name":"updated"}`, userBeforeUpdate.ID)

	req, _ := http.NewRequest("PUT", ts.URL+"/users/"+strconv.Itoa(userBeforeUpdate.ID),
		strings.NewReader(jsonUpdate))
	resp, err = http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	userUpdated := new(User)
	err = json.NewDecoder(resp.Body).Decode(userUpdated)
	assert.NoError(err)
	assert.Equal(userUpdated.FirstName, "updated")
	assert.Equal(userUpdated.LastName, "Han")
}
