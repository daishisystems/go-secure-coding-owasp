package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/add", addHandler)
	http.ListenAndServe(":8080", nil)
}

func addHandler(w http.ResponseWriter, r *http.Request) {
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
