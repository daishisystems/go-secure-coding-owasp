package main

import (
	"bytes"
	"fmt"
	"log"
)

func main() {
	var buf bytes.Buffer
	var RoleLevel int

	logger := log.New(&buf, "logger: ", log.Lshortfile)

	fmt.Println("Please enter your user level.")
	fmt.Scanf("%d", &RoleLevel) //<--- example

	switch RoleLevel {
	case 1:
		// Log successful login
		logger.Printf("Login successful.")
		fmt.Print(&buf)
	case 2:
		// Log unsuccessful Login
		logger.Printf("Login unsuccessful - Insufficient access level.")
		fmt.Print(&buf)
	default:
		// Unspecified error
		logger.Print("Login error.")
		fmt.Print(&buf)
	}
}
