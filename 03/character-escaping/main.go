package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/daishisystems/go-secure-coding-owasp/03/character-escaping/pkg/char"
)

type Handler struct {
	char *char.Char
}

func (h *Handler) AddHandler(w http.ResponseWriter, r *http.Request) {

	text := r.URL.Query().Get("text")
	if text == "" {
		http.Error(w, "No http request param matching 'text''", http.StatusBadRequest)
		return
	}

	result := h.char.Escape(text)
	payload := make(map[string]string)
	payload["escapedText"] = result

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
		char: char.Newchar(),
	}

	http.HandleFunc("/escape", handler.AddHandler)
	fmt.Println("Listening on port 8082")
	http.ListenAndServe(":8082", nil)
}
