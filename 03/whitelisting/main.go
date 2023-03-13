package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Define a map of allowed values
	allowedValues := map[string]bool{
		"foo": true,
		"bar": true,
		"baz": true,
	}

	// Define the HTTP handler function
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Get user input from query parameter
		userInput := r.URL.Query().Get("input")

		// Check if user input is in the map of allowed values
		if _, ok := allowedValues[userInput]; ok {
			fmt.Fprintf(w, "Valid input: %s", userInput)
		} else {
			fmt.Fprint(w, "Invalid input")
		}
	})

	// Start the web server
	fmt.Println("Listening on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
