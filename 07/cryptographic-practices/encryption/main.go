package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"log"
)

func encrypt(val []byte, secret []byte) ([]byte, error) {
	block, err := aes.NewCipher(secret)
	if err != nil {
		return nil, err
	}

	aead, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, aead.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return aead.Seal(nonce, nonce, val, nil), nil
}

func decrypt(val []byte, secret []byte) ([]byte, error) {
	block, err := aes.NewCipher(secret)
	if err != nil {
		return nil, err
	}

	aead, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	size := aead.NonceSize()
	if len(val) < size {
		return nil, err
	}

	result, err := aead.Open(nil, val[:size], val[size:], nil)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func secret() ([]byte, error) {
	key := make([]byte, 16)

	if _, err := rand.Read(key); err != nil {
		return nil, err
	}

	return key, nil
}

func main() {
	secret, err := secret()
	if err != nil {
		log.Fatalf("unable to create secret key: %v", err)
	}

	message := []byte("Welcome to Go Language Secure Coding Practices")
	log.Printf("Message  : %s\n", message)

	encrypted, err := encrypt(message, secret)
	if err != nil {
		log.Fatalf("unable to encrypt the data: %v", err)
	}
	log.Printf("Encrypted: %x\n", encrypted)

	decrypted, err := decrypt(encrypted, secret)
	if err != nil {
		log.Fatalf("unable to decrypt the data: %v", err)
	}
	log.Printf("Decrypted: %s\n", decrypted)
}
