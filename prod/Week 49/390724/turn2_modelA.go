package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// Define the secret key and other constants
const (
	secretKey     = "my_secret"
	tokenExpire   = time.Minute * 15 // Token expiration time
	refreshExpire = time.Hour * 24   // Refresh token expiration time
)

// Create a custom claims type
type CustomClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// GenerateJWT creates a new JWT token
func GenerateJWT(username string) (string, error) {
	claims := CustomClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "my_app",
			Subject:   username,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenExpire)),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

// ValidateJWT validates the JWT token
func ValidateJWT(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method and return the key
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		return nil, fmt.Errorf("invalid claims type")
	}

	// Validating the issued at and expiration claims
	if claims.IssuedAt.Time.After(time.Now()) {
		return nil, fmt.Errorf("token not yet valid")
	}
	if claims.ExpiresAt.Time.Before(time.Now()) {
		return nil, fmt.Errorf("token expired")
	}

	return claims, nil
}

// Handler to demonstrate JWT usage
func handler(w http.ResponseWriter, r *http.Request) {
	token, err := GenerateJWT("example_user")
	if err != nil {
		http.Error(w, "Could not generate token", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Generated Token: %s\n", token)

	// Validate the token
	claims, err := ValidateJWT(token)
	if err != nil {
		http.Error(w, "Token validation failed", http.StatusUnauthorized)
		return
	}

	fmt.Fprintf(w, "Token Validated! Welcome %s\n", claims.Username)
}

func main() {
	http.HandleFunc("/generate", handler)
	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
