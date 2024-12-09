package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

func generateHighSensitivityRSAKey() (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, 4096)
}

func main() {
	// RSA example for high sensitivity
	key, err := generateHighSensitivityRSAKey()
	if err != nil {
		fmt.Println("Error generating RSA key:", err)
		return
	}

	privateKeyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)})
	publicKeyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PUBLIC KEY", Bytes: x509.MarshalPKCS1PublicKey(&key.PublicKey)})

	fmt.Println("High Sensitivity RSA Private Key:\n", string(privateKeyPEM))
	fmt.Println("High Sensitivity RSA Public Key:\n", string(publicKeyPEM))
}
