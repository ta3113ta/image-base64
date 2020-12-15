package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"mime/multipart"
	"io"
	"path/filepath"
)

func TestUploadFile(t *testing.T) {
	// Read test image
	Path := "./test.png"
	file, err := os.Open(Path)
	if err != nil {
		t.Error(err)
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("my_file", filepath.Base(Path))
	if err != nil {
		writer.Close()
		t.Error(err)
	}
	io.Copy(part, file)
	writer.Close()

	// New request
	req := httptest.NewRequest("POST", "/upload", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res := httptest.NewRecorder()

	// When upload file
	uploadFile(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("Status code must be 200, your status code is : %d", res.Code)
	}
}