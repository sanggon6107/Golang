package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestFileUpload(t *testing.T) {
	assert := assert.New()
	path := ""
	file, _ := os.Open(path)
	defer file.Close()

	buf := &bytes.Buffer{}
	// MIME : Multipurpose Internet Mail Extensions is an Internet standard that extends
	// the format of email messages to support text in character sets other than ASCII.

	writer := multipart.NewWriter(buf)
	multi, err := writer.CreateFormFile("upload_file_name", filepath.Base(path))
	assert.NoError(err) // should not be an error.
	io.Copy(multi, file)
	writer.Close()

	res := httptest.NewRecoder()
	req := httptest.NewRequest("POST", "/uploads", buf)

	// The Content-Type header is used in web requests to indicate what type of media or resource is being used in the request or response.
	// When you send data in a request such as PUT(Update) or POST(Create),
	// you pass the Content-Type header to tell the server what type of data it is receiving.

	req.Header.set("Content-type", writer.FormDataContentType())
	FileUploadHandler(res, req)
	assert.Equal(http.StatusOK, res.Code)
}
