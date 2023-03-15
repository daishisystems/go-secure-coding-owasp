package main

import (
	"fmt"
	"net/http"

	"github.com/daishisystems/go-secure-coding-owasp/03/boundary-checking/pkg/calculator"
	"github.com/daishisystems/go-secure-coding-owasp/03/boundary-checking/pkg/validator"
)

type Handler struct {
	calculate       *calculator.Calculate
	integerValidate *validator.IntegerValidate
}

func (h *Handler) AddHandler(w http.ResponseWriter, r *http.Request) {
	num1Str := r.URL.Query().Get("num1")
	num2Str := r.URL.Query().Get("num2")

	num1, num2, err := h.integerValidate.IsValidInteger(num1Str, num2Str)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := h.calculate.Add(num1, num2)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "%d", result)
}

func main() {
	calculate := &calculator.Calculate{}
	integerValidate := &validator.IntegerValidate{}
	handler := &Handler{
		calculate:       calculate,
		integerValidate: integerValidate,
	}

	http.HandleFunc("/add", handler.AddHandler)
	http.ListenAndServe(":8080", nil)
}
