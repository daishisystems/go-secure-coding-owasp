package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"html/template"
	"image/png"
	"log"
	"net/http"

	"github.com/pquerna/otp/totp"
)

type TemplateData struct {
	Username    string
	QRCodeImage string
	SecretKey   string
}

var userSecretKeys map[string]string

func main() {
	userSecretKeys = make(map[string]string)
	http.HandleFunc("/", loginHandler)
	http.HandleFunc("/verify", verifyHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		username := r.FormValue("username")

		// Check if the user already has a secret key stored
		secretKey, ok := userSecretKeys[username]
		if !ok {
			// Generate TOTP key
			key, err := totp.Generate(totp.GenerateOpts{
				Issuer:      "globomantics.com",
				AccountName: username,
			})
			if err != nil {
				log.Fatal(err)
			}

			// Store the secret key
			userSecretKeys[username] = key.Secret()
			secretKey = key.Secret()

			// Convert TOTP key into a PNG image
			img, err := key.Image(200, 200)
			if err != nil {
				log.Fatal(err)
			}

			// Encode PNG image to byte slice
			var buf bytes.Buffer
			err = png.Encode(&buf, img)
			if err != nil {
				log.Fatal(err)
			}

			qrCodeImage := base64.StdEncoding.EncodeToString(buf.Bytes())

			// Prepare the data for the template
			data := TemplateData{
				Username:    username,
				QRCodeImage: qrCodeImage,
				SecretKey:   secretKey,
			}

			tmpl, err := template.ParseFiles("verify.html")
			if err != nil {
				log.Fatal(err)
			}

			err = tmpl.Execute(w, data)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			// User already has a secret key, display the verification form directly
			data := TemplateData{
				Username:  username,
				SecretKey: secretKey,
			}

			tmpl, err := template.ParseFiles("verify.html")
			if err != nil {
				log.Fatal(err)
			}

			err = tmpl.Execute(w, data)
			if err != nil {
				log.Fatal(err)
			}
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
		// Retrieve the passcode and secret key from the form
		passcode := r.FormValue("passcode")
		secretKey := r.FormValue("secretKey")

		// Validate the TOTP passcode
		valid := totp.Validate(passcode, secretKey)
		if valid {
			fmt.Fprintf(w, "Authentication successful!")
		} else {
			fmt.Fprintf(w, "Authentication failed!")
		}
	}
}
