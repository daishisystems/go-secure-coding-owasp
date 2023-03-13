package main

import (
	"fmt"
)

func main() {
	// Define a list of allowed values
	allowedValues := map[string]bool{
		"foo": true,
		"bar": true,
		"baz": true,
	}

	// Get user input
	var userInput string
	fmt.Println("Enter a value:")
	fmt.Scanln(&userInput)

	// Check if user input is in the list of allowed values
	if _, ok := allowedValues[userInput]; ok {
		fmt.Println("Valid input.")
	} else {
		fmt.Println("Invalid input.")
	}
}
