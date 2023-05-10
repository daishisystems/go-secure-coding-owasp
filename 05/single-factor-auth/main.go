package main

import (
	"database/sql"  // Package for working with SQL databases
	"fmt"           // Package to format text
	"html/template" // Package to parse and execute HTML templates
	"log"           // Package for logging
	"net/http"      // Package for HTTP requests and responses

	_ "github.com/go-sql-driver/mysql" // MySQL driver
	"golang.org/x/crypto/bcrypt"       // Package to hash and compare passwords
)

// Define a struct to hold information about a user
type user struct {
	id       int
	username string
	password string
}

func main() {
	// Connect to MySQL database using a DSN (Data Source Name)
	db, err := sql.Open("mysql", "root:admin@tcp(localhost:3306)/globomantics")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close() // Defer closing the database connection until the end of the function

	// Serve the login page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			// Get the username and password submitted in the login form
			username := r.FormValue("username")
			password := r.FormValue("password")

			// Query the database for the user with the given username
			row := db.QueryRow("SELECT id, username, password FROM users WHERE username = ?", username)
			var u user
			err := row.Scan(&u.id, &u.username, &u.password)
			if err != nil {
				// Log the error and return an HTTP 401 Unauthorized response
				log.Printf("Failed login attempt for user %q: %v", username, err)
				http.Error(w, "Invalid username or password", http.StatusUnauthorized)
				return
			}

			// Check if the password is correct using bcrypt
			err = bcrypt.CompareHashAndPassword([]byte(u.password), []byte(password))
			if err != nil {
				// Log the error and return an HTTP 401 Unauthorized response
				log.Printf("Failed login attempt for user %q: %v", username, err)
				http.Error(w, "Invalid username or password", http.StatusUnauthorized)
				return
			}

			// If the login was successful, display a welcome message
			fmt.Fprintf(w, "Welcome, %s!", u.username)
			return
		}

		// If the login form was not submitted yet, serve the login form
		tmpl, err := template.ParseFiles("login.html") // Parse the login form template
		if err != nil {
			log.Fatal(err)
		}

		// Set the Content-Type header of the response to "text/html"
		w.Header().Set("Content-Type", "text/html")

		// Execute the login form template with no data and write the result to the response
		err = tmpl.Execute(w, nil)
		if err != nil {
			log.Fatal(err)
		}
	})

	// Listen on port 8080 and serve requests indefinitely
	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
