package main

import (
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Define a secret key for signing the tokens
var jwtKey = []byte("my_secret_key")

// Create a struct for the claims
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenerateJWT generates a new JWT token with an expiration time
func GenerateJWT(username string, expirationTime time.Duration) (string, error) {
	expiration := time.Now().Add(expirationTime) // Set token expiration time
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiration.Unix(), // Set expiration time in seconds
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey) // Sign the token
}

// ValidateJWT validates the token and checks its expiration
func ValidateJWT(tokenString string) error {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	// Check if the token has expired
	if claims.ExpiresAt < time.Now().Unix() {
		return fmt.Errorf("token expired")
	}

	return nil
}

func main() {
	// Generate a token that expires in 5 minutes
	token, err := GenerateJWT("user1", 5*time.Minute)
	if err != nil {
		log.Fatalf("Error generating token: %v", err)
	}

	fmt.Println("Generated Token:", token)

	// Validate the token
	err = ValidateJWT(token)
	if err != nil {
		log.Fatalf("Token validation failed: %v", err)
	} else {
		fmt.Println("Token is valid!")
	}

	// Wait for 6 minutes to test expiration
	time.Sleep(6 * time.Minute)

	// Validate the token again
	err = ValidateJWT(token)
	if err != nil {
		log.Printf("Token validation failed after expiration: %v", err)
	} else {
		fmt.Println("Token is valid!")
	}
}
