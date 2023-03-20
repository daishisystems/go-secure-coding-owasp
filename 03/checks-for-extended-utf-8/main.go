package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/daishisystems/go-secure-coding-owasp/03/checks-for-extended-utf-8/pkg/utf"
)

type Handler struct {
	utf *utf.Utf
}

func (h *Handler) AddHandler(w http.ResponseWriter, r *http.Request) {

	input := r.URL.Query().Get("input")
	if input == "" {
		http.Error(w, "No http request param matching 'input'", http.StatusBadRequest)
		return
	}

	inputIsValudUtf8 := h.utf.InputIsValidUtf8(input)

	payload := make(map[string]bool)
	payload["inputIsValudUtf8"] = inputIsValudUtf8

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
		utf: utf.NewUtf(),
	}

	http.HandleFunc("/inputIsValudUtf8", handler.AddHandler)
	fmt.Println("Listening on port 8087")
	http.ListenAndServe(":8087", nil)
}
