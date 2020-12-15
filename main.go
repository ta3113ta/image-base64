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
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving file from form-data")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// Read to byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	encoded := base64.StdEncoding.EncodeToString(fileBytes)
	if len(encoded) == 0 {
		fmt.Println("eiei")
	}
	// return whether or not this has been successful
	fmt.Fprintf(w, "Successfully Uploaded File\n")

}

func setupRoutes() {
	http.HandleFunc("/upload", uploadFile)
	http.ListenAndServe(":8080", nil)
}

func main() {
	fmt.Println("Go File Upload Tutorial")
	setupRoutes()
}