package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"slices"
	"time"
)

func measureAESPerformance(keyLength int) {
	// Valid key lengths for AES are 128, 192, and 256 bits

	if !slices.Contains([]int{128, 192, 256}, keyLength) {
		fmt.Println("Invalid key length. Must be 128, 192, or 256.")
		return
	}

	// Generate a random key of the specified length
	key := make([]byte, keyLength/8)
	if _, err := rand.Read(key); err != nil {
		fmt.Println("Error generating key:", err)
		return
	}

	// Generate a random plaintext
	plaintext := make([]byte, 1024*1024) // 1MB plaintext
	if _, err := rand.Read(plaintext); err != nil {
		fmt.Println("Error generating plaintext:", err)
		return
	}

	// Create an AES block cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("Error creating AES cipher:", err)
		return
	}

	// Ciphertext size must be plaintext size + block size.
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		fmt.Println("Error generating IV:", err)
		return
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	// Measure encryption time
	start := time.Now()
	for i := 0; i < 100; i++ { // Repeat 100 times for better averaging
		stream = cipher.NewCFBEncrypter(block, iv)
		stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)
	}
	end := time.Now()

	fmt.Printf("AES-%d bit encryption time: %v\n", keyLength, end.Sub(start))
}

func main() {
	measureAESPerformance(128)
	measureAESPerformance(192)
	measureAESPerformance(256)
}
