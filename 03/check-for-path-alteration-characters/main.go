package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/daishisystems/go-secure-coding-owasp/03/check-for-path-alteration-characters/pkg/path"
)

type Handler struct {
	path *path.Path
}

func (h *Handler) AddHandler(w http.ResponseWriter, r *http.Request) {

	pathIsValid := h.path.PathIsValid(r.RequestURI)

	payload := make(map[string]bool)
	payload["pathIsValid"] = pathIsValid

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
		path: path.NewPath(),
	}

	http.HandleFunc("/pathIsValid", handler.AddHandler)
	fmt.Println("Listening on port 8086")
	http.ListenAndServe(":8086", nil)
}
