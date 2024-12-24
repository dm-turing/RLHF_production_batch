package main

import (
	"fmt"
	"os"
)

func getEncryptionKey() string {
	key := os.Getenv("ENCRYPTION_KEY")
	if key == "" {
		panic("Encryption key not found in environment variables.")
	}
	return key
}

func main() {
	fmt.Println(getEncryptionKey())
}
