package main

import (
	"fmt"
	"net/http"

	"github.com/daishisystems/go-secure-coding-owasp/03/boundary-checking/pkg/calculator"
	"github.com/daishisystems/go-secure-coding-owasp/03/boundary-checking/pkg/validator"
)

type Handler struct {
	calculator       *calculator.Calculator
	integerValidator *validator.IntegerValidator
}

func (h *Handler) AddHandler(w http.ResponseWriter, r *http.Request) {
	num1Str := r.URL.Query().Get("num1")
	num2Str := r.URL.Query().Get("num2")

	num1, err := h.integerValidator.Validate(num1Str)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	num2, err := h.integerValidator.Validate(num2Str)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := h.calculator.Add(num1, num2)
	fmt.Fprintf(w, "%d", result)
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
