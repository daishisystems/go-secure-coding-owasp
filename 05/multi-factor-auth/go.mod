module github.com/daishisystems/go-secure-coding-owasp/05/single-factor-auth

go 1.20

require (
	github.com/go-sql-driver/mysql v1.7.1
	github.com/gorilla/sessions v1.2.1
	golang.org/x/crypto v0.8.0
)

require github.com/boombuler/barcode v1.0.1-0.20190219062509-6c824513bacc // indirect

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gorilla/securecookie v1.1.1 // indirect
	github.com/pquerna/otp v1.4.0
)
