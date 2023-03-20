package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/daishisystems/go-secure-coding-owasp/03/check-for-null-bytes/pkg/validate"
)

type Handler struct {
	validator *validate.Validator
}

func (h *Handler) AddHandler(w http.ResponseWriter, r *http.Request) {

	input := r.URL.Query().Get("input")
	if input == "" {
		http.Error(w, "No http request param matching 'input'", http.StatusBadRequest)
		return
	}

	containsNullBytes := h.validator.ContainsNullByte(input)

	payload := make(map[string]bool)
	payload["containsNullByte"] = containsNullBytes

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
		validator: validate.NewValidator(),
	}

	http.HandleFunc("/containsnullbyte", handler.AddHandler)
	fmt.Println("Listening on port 8084")
	http.ListenAndServe(":8084", nil)
}
