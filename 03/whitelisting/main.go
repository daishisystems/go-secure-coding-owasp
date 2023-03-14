package main

import (
	"fmt"
	"net/http"

	inputHandler "github.com/daishisystems/go-secure-coding-owasp/03/whitelisting/pkg/handlers/input"
	inputService "github.com/daishisystems/go-secure-coding-owasp/03/whitelisting/pkg/services/input"
)

func main() {
	inputValues := inputService.NewInputValues()
	inputHandler := inputHandler.NewInputHandler(inputValues)

	server := http.Server{
		Addr:    ":8080",
		Handler: inputHandler,
	}

	fmt.Println("Listening on port 8080")
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
