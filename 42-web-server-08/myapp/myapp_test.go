package myapp

import (
	"github.com/stretchr/testify/assert"
	"ioutil"
	"net/http/httptest"
	"testing"
)

func TestNewHandler(t *testing.T) {
	assert := assert.New()
	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	resp, err := http.Get(ts.URL) // get response by "Get"
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	data, _ := ioutil.ReadAll(resp.Body)
	assert.Equal("Hello world!", string(data))
}

func TestUsers(t *testing.T) {
	assert := assert.New()
	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/users")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	data, _ := ioutil.ReadAll(resp.Body) // response
	assert.Contains(string(data), "Get UserInfo")
}
