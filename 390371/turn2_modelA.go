package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

// Sensitivity levels
const (
	PublicData          = "public"
	ConfidentialData    = "confidential"
	HighlySensitiveData = "highly_sensitive"
)

// SelectKeyLength determines the appropriate key length based on data sensitivity.
func SelectKeyLength(sensitivity string) int {
	switch sensitivity {
	case PublicData:
		return 16 // 128 bits
	case ConfidentialData:
		return 32 // 256 bits
	case HighlySensitiveData:
		return 48 // 384 bits, custom for example; AES max is actually 32 bytes (AES-256)
	default:
		return 32 // default to 256 bits
	}
}

// Create a new AES cipher block
func NewCipher(key []byte) (cipher.Block, error) {
	return aes.NewCipher(key)
}

// Encrypt message using AES
func Encrypt(key []byte, plaintext string) (string, error) {
	block, err := NewCipher(key)
	if err != nil {
		return "", err
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], []byte(plaintext))

	return hex.EncodeToString(ciphertext), nil
}

// Decrypt message using AES
func Decrypt(key []byte, ciphertext string) (string, error) {
	ciphertextBytes, _ := hex.DecodeString(ciphertext)
	block, err := NewCipher(key)
	if err != nil {
		return "", err
	}

	if len(ciphertextBytes) < aes.BlockSize {
		return "", fmt.Errorf("ciphertext too short")
	}

	iv := ciphertextBytes[:aes.BlockSize]
	ciphertextBytes = ciphertextBytes[aes.BlockSize:]

	mode := cipher.NewCBCDecrypter(block, iv)

	mode.CryptBlocks(ciphertextBytes, ciphertextBytes)

	return string(ciphertextBytes), nil
}

// Main function to demonstrate adaptive key usage
func main() {
	sensitivity := ConfidentialData // Change sensitivity levels here

	keyLength := SelectKeyLength(sensitivity)

	// Generate a 32, 48, or 16-byte key based on the key length for AES
	key := make([]byte, keyLength)
	if _, err := rand.Read(key); err != nil {
		fmt.Println("Error generating key:", err)
		return
	}

	plaintext := "This is a secret message by user"
	fmt.Printf("Sensitive Data: %s\n", sensitivity)
	fmt.Printf("Using Key Length: %d bytes\n", keyLength)

	// Encrypt the message
	ciphertext, err := Encrypt(key, plaintext)
	if err != nil {
		fmt.Println("Error encrypting:", err)
		return
	}

	fmt.Printf("Encrypted Message: %s\n", ciphertext)

	// Decrypt the message
	decrypted, err := Decrypt(key, ciphertext)
	if err != nil {
		fmt.Println("Error decrypting:", err)
		return
	}

	fmt.Printf("Decrypted Message: %s\n", decrypted)
}
