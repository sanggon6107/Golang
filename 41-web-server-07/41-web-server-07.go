package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func FileUploadHandler(w http.ResponseWriter, r *http.Request) {
	uploadedFile, header, err := r.FormFile("upload_file_name")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	defer uploadedFile.Close()

	pathUploadedFile := "./uploads"
	os.MkdirAll(pathUploadedFile, 0777) // mkdir if the path doesn't exist
	filePath := fmt.Sprintf("%s/%s", pathUploadedFile, header.Filename)
	createdFile, err := os.Create(filePath)
	defer createdFile.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}
	io.Copy(createdFile, uploadedFile)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, filePath)
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("public")))

	http.HandleFunc("/uploads", FileUploadHandler)

	http.ListenAndServe(":3000", nil)
}
