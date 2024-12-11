package main

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func main() {
	// Example: Generating a signed JWT with claims
	secret := []byte("super-secret-key") // Replace with a secure key
	iat := time.Now().Unix() - 61
	exp := iat + 3600 // Token expires in 1 hour
	nbf := iat + 60   // Token is not before 1 minute
	// jti := jwt.NewUUID() // Unique identifier for the token

	claims := &jwt.StandardClaims{
		Audience:  "your-audience",
		Issuer:    "your-issuer",
		Subject:   "1234567890",
		ExpiresAt: exp,
		NotBefore: nbf,
		IssuedAt:  iat,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secret)
	if err != nil {
		fmt.Println("Error signing token:", err)
		return
	}

	fmt.Println("Signed Token:", signedToken)

	// Example: Validating the JWT
	token, err = jwt.ParseWithClaims(signedToken, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})

	if err != nil {
		fmt.Println("Error parsing token:", err)
		return
	}

	claims = token.Claims.(*jwt.StandardClaims)
	err = claims.Valid()
	if err != nil {
		fmt.Println("Token is invalid")
		return
	}

	fmt.Println("Token is valid:")
	fmt.Println("  Issuer:", claims.Issuer)
	fmt.Println("  Subject:", claims.Subject)
	fmt.Println("  Audience:", claims.Audience)
	fmt.Println("  ExpiresAt:", time.Unix(claims.ExpiresAt, 0))
	fmt.Println("  NotBefore:", time.Unix(claims.NotBefore, 0))
	fmt.Println("  IssuedAt:", time.Unix(claims.IssuedAt, 0))
	// fmt.Println("  JTI:", claims.JTI)
}
