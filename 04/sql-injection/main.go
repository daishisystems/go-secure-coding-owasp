package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/daishisystems/go-secure-coding-owasp/04/sql-injection/pkg/patient"
)

func main() {
	http.HandleFunc("/search", patient.HandleSearch)
	fmt.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
