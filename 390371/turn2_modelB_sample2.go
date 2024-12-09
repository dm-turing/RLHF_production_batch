package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io"
)

func encryptModerateSensitivity(data []byte) ([]byte, error) {
	// AES-256 is used for moderate sensitivity data
	block, err := aes.NewCipher([]byte("12345678901234567890123456789012")) // Replace with a secure key
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

func generateRSAKey() (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, 2048)
}

func main() {
	// AES example
	plaintext := []byte("This is moderate sensitivity data.")
	ciphertext, err := encryptModerateSensitivity(plaintext)
	if err != nil {
		fmt.Println("Error encrypting:", err)
		return
	}
	fmt.Println("Ciphertext (AES-256):", ciphertext)

	// RSA example
	key, err := generateRSAKey()
	if err != nil {
		fmt.Println("Error generating RSA key:", err)
		return
	}

	privateKeyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)})
	publicKeyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PUBLIC KEY", Bytes: x509.MarshalPKCS1PublicKey(&key.PublicKey)})

	fmt.Println("RSA Private Key:\n", string(privateKeyPEM))
	fmt.Println("RSA Public Key:\n", string(publicKeyPEM))
}
