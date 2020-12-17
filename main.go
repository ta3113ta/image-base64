package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
	// Maximum size 10 MB.
	r.ParseMultipartForm(10 << 20)

	// retrieve file from posted form-data
	file, _, err := r.FormFile("inpFile")
	if err != nil {
		fmt.Println("Error Retrieving file from form-data")
		fmt.Println(err)
		return
	}
	defer file.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	encoded := base64.StdEncoding.EncodeToString(fileBytes)
	fmt.Fprintf(w, encoded)
}

func setupRoutes() {
	http.HandleFunc("/upload", uploadFile)
	http.ListenAndServe(":8080", nil)
}

func main() {
	fmt.Println("Go File Upload Tutorial")
	setupRoutes()
}