package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
)

// Helper function to randomize 128-bit key
func randKey(size int) []byte {
	key := make([]byte, size)
	/*_, err := io.ReadFull(rand.Reader, key)
	if err != nil {
		fmt.Println("Error reading bytes:", err)
	}*/
	return key
}

func encrypt(data []byte, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	ciphertext := make([]byte, aes.BlockSize+len(data))
	iv := ciphertext[:aes.BlockSize]
	_, err = io.ReadFull(rand.Reader, iv)
	if err != nil {
		return "", err
	}

	cbc := cipher.NewCBCEncrypter(block, iv)
	cbc.CryptBlocks(ciphertext[aes.BlockSize:], data)

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func decrypt(ciphertext string, key []byte) ([]byte, error) {
	decoded, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(decoded) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}

	iv := decoded[:aes.BlockSize]
	ciphertextBytes := decoded[aes.BlockSize:]

	cbc := cipher.NewCBCDecrypter(block, iv)
	cbc.CryptBlocks(ciphertextBytes, ciphertextBytes)

	return ciphertextBytes, nil
}

func main() {
	key := randKey(32) // 128-bit key
	data := []byte("This is a secret message")
	enc, err := encrypt(data, key)
	if err != nil {
		fmt.Println("Encryption error:", err)
		return
	}
	fmt.Println("Encrypted data:", enc)

	dec, err := decrypt(enc, key)
	if err != nil {
		fmt.Println("Decryption error:", err)
		return
	}
	fmt.Println("Decrypted data:", string(dec))
}
