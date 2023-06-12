package main

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"io"

	"golang.org/x/crypto/blake2s"
)

func main() {
	md5Hash := md5.New()
	sha256Hash := sha256.New()
	h_blake2s, _ := blake2s.New256(nil)
	io.WriteString(md5Hash, "Welcome to Go Language Secure Coding Practices")
	io.WriteString(sha256Hash, "Welcome to Go Language Secure Coding Practices")
	io.WriteString(h_blake2s, "Welcome to Go Language Secure Coding Practices")
	fmt.Printf("MD5        : %x\n", md5Hash.Sum(nil))
	fmt.Printf("SHA256     : %x\n", sha256Hash.Sum(nil))
	fmt.Printf("Blake2s-256: %x\n", h_blake2s.Sum(nil))
}
