package main

import (
	"fmt"
	"net/http"
	"os"
)

const (
	logoFile       = "./logo.jpg"
	maxContentType = 512
)

var allowedFileTypes = []string{
	"image/jpeg",
	"image/jpg",
	"image/gif",
	"image/png",
}

func main() {
	file, err := os.Open(logoFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	buff := make([]byte, maxContentType)
	_, err = file.Read(buff)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	filetype := http.DetectContentType(buff)

	if isValidFileType(filetype) {
		fmt.Println(filetype)
	} else {
		fmt.Println("unknown file type uploaded")
	}
}

func isValidFileType(filetype string) bool {
	for _, allowedType := range allowedFileTypes {
		if filetype == allowedType {
			return true
		}
	}
	return false
}
