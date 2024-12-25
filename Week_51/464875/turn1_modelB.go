package main

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Set your secret key (should be securely managed)
var jwtSecretKey = []byte("SecretPassphrase")

// Define the claims structure
type CustomClaims struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	jwt.StandardClaims
}

func generateToken(userID string, role string) (string, error) {
	// Create new claims
	claims := CustomClaims{
		UserID: userID,
		Role:   role,
		StandardClaims: jwt.StandardClaims{
			// Set token expiration time here (e.g., 2 hours)
			ExpiresAt: time.Now().Add(2 * time.Hour).Unix(),
		},
	}

	// Create the JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecretKey)
	return tokenString, err
}

func validateToken(tokenString string) (*CustomClaims, error) {
	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Check for the correct signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// Return the secret key to validate the signature
		return jwtSecretKey, nil
	})

	if err != nil {
		return nil, err
	}

	// Assert the claims is of the correct type
	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		return nil, fmt.Errorf("Can't decode claims")
	}

	return claims, nil
}

func main() {
	token, _ := generateToken("user", "admin")
	c, _ := validateToken(token)
	fmt.Println(c.UserID, c.Role)
}
