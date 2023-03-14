package input

import (
	"fmt"
	"net/http"

	"github.com/daishisystems/go-secure-coding-owasp/03/whitelisting/pkg/services/input"
)

type InputHandler struct {
	service input.InputService
}

func NewInputHandler(service input.InputService) *InputHandler {
	return &InputHandler{
		service: service,
	}
}

func (ih *InputHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	userInput := r.URL.Query().Get("input")

	if ih.service.CheckInput(userInput) {
		fmt.Fprintf(w, "Valid input: %s", userInput)
	} else {
		fmt.Fprint(w, "Invalid input")
	}
}
