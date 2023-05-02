package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		input := r.URL.Query().Get("input")
		output := "<h1>Hello, " + input + "!</h1>"
		io.WriteString(w, output)
	})

	http.ListenAndServe(":8080", nil)
}
