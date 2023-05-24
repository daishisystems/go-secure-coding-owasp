package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/pquerna/otp/totp"
)

const (
	secretKey = "VIZHFQWE2M4OEOGK"
)

func main() {
	http.HandleFunc("/", loginHandler)
	http.HandleFunc("/verify", verifyHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")

		// Validate username and password against your user database
		if isValidCredentials(username, password) {
			// Generate TOTP for the user
			totpCode := generateTOTP(username)

			// Render the TOTP verification page
			tmpl, err := template.ParseFiles("verify.html")
			if err != nil {
				log.Fatal(err)
			}

			data := struct {
				Username string
			}{
				Username: username,
			}

			err = tmpl.Execute(w, data)
			if err != nil {
				log.Fatal(err)
			}

			// Send the TOTP code to the user (e.g., via email or SMS)
			fmt.Println("TOTP Code:", totpCode)
		} else {
			// Invalid credentials, show error message
			fmt.Fprintf(w, "Invalid credentials")
		}
	} else {
		// Render the login page
		tmpl, err := template.ParseFiles("login.html")
		if err != nil {
			log.Fatal(err)
		}

		err = tmpl.Execute(w, nil)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func verifyHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		username := r.FormValue("username")
		totpCode := r.FormValue("totp")

		// Verify the entered TOTP code
		if verifyTOTP(username, totpCode) {
			fmt.Fprintf(w, "Authentication successful!")
		} else {
			fmt.Fprintf(w, "Authentication failed!")
		}
	}
}

func isValidCredentials(username, password string) bool {
	// Implement your logic to validate username and password
	// against your user database. Return true if valid, false otherwise.
	return true
}

func generateTOTP(username string) string {
	// Generate a TOTP code for the user using the secret key
	key := []byte(secretKey + username)
	totpCode, err := totp.GenerateCode(string(key), time.Now())
	if err != nil {
		log.Fatal(err)
	}

	return totpCode
}

func verifyTOTP(username, totpCode string) bool {
	// Verify the entered TOTP code against the secret key
	key := []byte(secretKey + username)
	return totp.Validate(totpCode, string(key))
}
