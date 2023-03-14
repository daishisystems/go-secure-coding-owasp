package main

import (
	"net/http"

	"github.com/daishisystems/go-secure-coding-owasp/03/boundary-checking/handlers/addhandler"
)

func main() {
	addHandler := addhandler.NewAddHandler()
	http.HandleFunc("/add", addHandler.Handle)
	http.ListenAndServe(":8080", nil)
}
