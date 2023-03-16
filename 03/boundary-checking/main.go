package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/daishisystems/go-secure-coding-owasp/03/boundary-checking/pkg/calculator"
	"github.com/daishisystems/go-secure-coding-owasp/03/boundary-checking/pkg/validator"
)

type Handler struct {
	calculator       *calculator.Calculator
	integerValidator *validator.IntegerValidator
}

// todo: Use pagination range example.
func (h *Handler) AddHandler(w http.ResponseWriter, r *http.Request) {
	num1Str := r.URL.Query().Get("num1")
	num1, err := h.integerValidator.Validate(num1Str)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	num2Str := r.URL.Query().Get("num2")
	num2, err := h.integerValidator.Validate(num2Str)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := h.calculator.Add(num1, num2)
	response := make(map[string]int)
	response["result"] = result

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	jsonResp, err := json.Marshal(response)
	if err != nil {
		log.Fatalf("An error occured while marshalling the response. Err: %s", err)
	}

	w.Write(jsonResp)
}

func main() {
	handler := &Handler{
		calculator:       calculator.NewCalculator(),
		integerValidator: validator.NewIntegerValidator(),
	}

	http.HandleFunc("/add", handler.AddHandler)
	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
