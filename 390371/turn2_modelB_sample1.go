package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

func encryptLowSensitivity(data []byte) ([]byte, error) {
	// AES-128 is used for low sensitivity data
	block, err := aes.NewCipher([]byte("1234567890123456")) // Replace with a secure key
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext, nil
}

func main() {
	plaintext := []byte("This is low sensitivity data.")
	ciphertext, err := encryptLowSensitivity(plaintext)
	if err != nil {
		fmt.Println("Error encrypting:", err)
		return
	}
	fmt.Println("Ciphertext:", ciphertext)
}
