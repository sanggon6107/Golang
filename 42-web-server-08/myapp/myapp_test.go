package myapp

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strconv"
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

func TestgetUsersHandler(t *testing.T) {
	assert := assert.New(t)
	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	userId := strconv.Itoa(rand.Intn(99))
	resp, err := http.Get(ts.URL + "/users/" + userId)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	data, _ := ioutil.ReadAll(resp.Body) // response
	assert.Contains(string(data), "User Id:66")
}
