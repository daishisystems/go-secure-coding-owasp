package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/daishisystems/go-secure-coding-owasp/03/whitelisting/pkg/whitelist"
)

type Handler struct {
	whitelist *whitelist.Whitelist
}

func (h *Handler) AddHandler(w http.ResponseWriter, r *http.Request) {

	input := r.URL.Query().Get("input")
	if input == "" {
		http.Error(w, "No http request param matching 'input''", http.StatusBadRequest)
		return
	}

	result := h.whitelist.Check(input)
	payload := make(map[string]bool)
	payload["isValid"] = result

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
		whitelist: whitelist.NewWhitelist(),
	}

	http.HandleFunc("/check", handler.AddHandler)
	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
