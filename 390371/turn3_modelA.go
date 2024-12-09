package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"time"
)

// Encrypts the given data using AES with the specified key length.
func encryptAes(key []byte, data []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
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

	return gcm.Seal(nonce, nonce, data, nil), nil
}

// Measure encryption time for different key lengths
func measureEncryptionTime(keySize int, data []byte) {
	key := make([]byte, keySize/8)
	if _, err := rand.Read(key); err != nil {
		fmt.Println("Error generating key:", err)
		return
	}

	start := time.Now()
	_, err := encryptAes(key, data)
	if err != nil {
		fmt.Println("Error encrypting:", err)
		return
	}
	elapsed := time.Since(start)
	fmt.Printf("Encryption with %d-bit key took %s\n", keySize, elapsed)
}

func main() {
	data := []byte("Sensitive data that needs to be encrypted securely.")

	// Measure performance for different key lengths
	measureEncryptionTime(128, data) // Low sensitivity
	measureEncryptionTime(192, data) // Moderate sensitivity
	measureEncryptionTime(256, data) // High sensitivity
}
