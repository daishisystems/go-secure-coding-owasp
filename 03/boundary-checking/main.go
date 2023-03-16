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
	if num1Str == "" {
		http.Error(w, "No http request param matching 'num1'", http.StatusBadRequest)
		return
	}

	num1, err := h.integerValidator.Validate(num1Str)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	num2Str := r.URL.Query().Get("num2")
	if num2Str == "" {
		http.Error(w, "No http request param matching 'num2'", http.StatusBadRequest)
		return
	}

	num2, err := h.integerValidator.Validate(num2Str)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sum := h.calculator.Add(num1, num2)
	payload := make(map[string]int)
	payload["sum"] = sum

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	body, err := json.Marshal(payload)
	if err != nil {
		log.Fatalf("An error occured while marshalling the response. Err: %s", err)
	}

	w.Write(body)
}

func main() {
	handler := &Handler{
		calculator:       calculator.NewCalculator(),
		integerValidator: validator.NewIntegerValidator(),
	}

	http.HandleFunc("/add", handler.AddHandler)
	fmt.Println("Listening on port 8081")
	http.ListenAndServe(":8081", nil)
}
