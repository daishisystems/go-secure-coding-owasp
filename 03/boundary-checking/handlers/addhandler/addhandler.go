package addhandler

import (
	"fmt"
	"net/http"
	"strconv"
)

type AddHandler struct {
}

func NewAddHandler() *AddHandler {
	return &AddHandler{}
}

func (h *AddHandler) Handle(w http.ResponseWriter, r *http.Request) {
	xStr := r.URL.Query().Get("start")
	x, err := strconv.Atoi(xStr)
	if err != nil {
		http.Error(w, "Invalid input for start", http.StatusBadRequest)
		return
	}

	if x < 0 || x > 100 {
		http.Error(w, "Input out of range for start", http.StatusBadRequest)
		return
	}

	yStr := r.URL.Query().Get("end")
	y, err := strconv.Atoi(yStr)
	if err != nil {
		http.Error(w, "Invalid input for end", http.StatusBadRequest)
		return
	}

	if y < 0 || y > 100 {
		http.Error(w, "Input out of range for end", http.StatusBadRequest)
		return
	}

	sum := x + y
	fmt.Fprintf(w, "Sum of %s and %s is %d", xStr, yStr, sum)
}
