package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// Define a struct to represent the JWT payload
type Payload struct {
	Subject string `json:"sub"`
	Name    string `json:"name"`
	Iat     int64  `json:"iat"` // Issued at
	Exp     int64  `json:"exp"` // Expiration
	Nbf     int64  `json:"nbf"` // Not before
}

func main() {
	// Create a new payload with the current time as issued at (iat)
	iat := time.Now().Unix()
	exp := iat + 3600 // Token expires in 1 hour
	nbf := iat + 60   // Token is not before 1 minute

	payload := Payload{
		Subject: "1234567890",
		Name:    "John Doe",
		Iat:     iat,
		Exp:     exp,
		Nbf:     nbf,
	}

	// Convert the payload to a JSON string
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error marshalling payload:", err)
		return
	}

	fmt.Println("Payload JSON:", string(payloadBytes))

	// Example: Verify if the token has expired
	currentTime := time.Now().Unix()
	if currentTime > exp {
		fmt.Println("Token has expired.")
	} else {
		fmt.Println("Token is still valid.")
	}

	// Example: Verify if the token is not before a certain time
	if currentTime < nbf {
		fmt.Println("Token is not yet valid.")
	} else {
		fmt.Println("Token is valid.")
	}
}
