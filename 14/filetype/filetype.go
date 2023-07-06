package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	file, err := os.Open("./logo.png")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	buff := make([]byte, 512)
	_, err = file.Read(buff)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	filetype := http.DetectContentType(buff)

	switch filetype {
	case "image/jpeg", "image/jpg", "image/gif", "image/png":
		fmt.Println(filetype)
	default:
		fmt.Println("unknown file type uploaded")
	}
}
