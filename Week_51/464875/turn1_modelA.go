package main

import (
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// FilePermissions represents the file permissions allowed for a user
type FilePermissions struct {
	UserID    string    `json:"user_id"`
	Files     []string  `json:"files"`
	AllowedAt time.Time `json:"allowed_at"`
}

func createJWT(userID string, filePermissions FilePermissions) (string, error) {
	// Define a private key
	privKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privateKeyPEM))
	if err != nil {
		return "", err
	}

	// Create a JWT claims struct
	claims := &jwt.StandardClaims{
		Audience:  "file-service",
		Issuer:    "auth-service",
		Subject:   userID,
		ExpiresAt: time.Now().Add(time.Hour).Unix(),
	}

	// Add custom claims
	claims.CustomClaims = filePermissions

	// Create a token
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	// Sign the token with the private key
	signedToken, err := token.SignedString(privKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func verifyJWT(tokenString string) (FilePermissions, error) {
	// Define a public key
	pubKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(publicKeyPEM))
	if err != nil {
		return FilePermissions{}, err
	}

	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return pubKey, nil
	})
	if err != nil {
		return FilePermissions{}, err
	}

	// Verify the token
	if !token.Valid {
		return FilePermissions{}, fmt.Errorf("Token is invalid")
	}

	var claims FilePermissions
	if err := token.Claims.(jwt.MapClaims).UnmarshalJSON([]byte(token.Claims.(jwt.MapClaims)["custom_claims"].(string)), &claims); err != nil {
		return FilePermissions{}, err
	}

	return claims, nil
}

func main() {
	// Create example claims
	claims := FilePermissions{
		UserID:    "user123",
		Files:     []string{"/path/to/file1", "/path/to/file2"},
		AllowedAt: time.Now(),
	}

	// Create a JWT
	tokenString, err := createJWT("user123", claims)
	if err != nil {
		log.Fatalf("Failed to create JWT: %v", err)
	}

	// Print JWT
	fmt.Println("JWT:", tokenString)

	// Verify the JWT
	verifiedClaims, err := verifyJWT(tokenString)
	if err != nil {
		log.Fatalf("Failed to verify JWT: %v", err)
	}

	// Print verified claims
	fmt.Println("Verified Claims:")
	fmt.Printf("  User ID: %s\n", verifiedClaims.UserID)
	fmt.Printf("  Files: %v\n", verifiedClaims.Files)
	fmt.Printf("  Allowed At: %s\n", verifiedClaims.AllowedAt)
}

// Sample RSA private and public keys
const privateKeyPEM = `
-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQDCyCWFLv+yWr7/UWOWucXEjv5cXHacb63n0Z0AgsACDk4kVe4/
Q2H9ZMrPww3iNgSIO8M02P4vjECcJrRbEGzFVzUux5c7OESpNGOw5iCLR20rI+X/
y01Vglv3ztRiHxMMskzTR3Btk4QSCNAPWuv1R2T6Z8Z/T4k9Xe6C5HprqQIDAQAB
AoGBAK//Lvq2K6qrA1Mw68/z/B501o9RctKkCe90KQpYyNnmRCxCITXMM7bmlyRT
TZewN3x0fJlPj0+yqIgGKkO4PlR3m7L3S+9h901SF3jUxzwESmU+LK16zCLvgK0v
3F1vc4JzZlF/KGfppfvuEQ3WXvixfvwmGmlInRfkHnUmAbZrAkEA6yFLNyvtV4b/
AxFv4SkCO3AVK82mTNTGLAzLxP/iIVeF6iS9aezWLDF/6TjFdsSNw5PG+lUg56bg
1NGf29HbIQJBAMCiBx2S20NV4OjNKEHDsbr0h8A5vVT1USMDlb2Xh38pG2AR5spT
d/NgICYr7Gp8PS0yzU9Jr/Dgq2S/UqNHb8MCQQDAMuPO34uGFyUNo3o7Q60Y1BG4
xYVTzJReUgfBRfCx8ZKZrNdxXG7oSmLVv1/MbELf2X63N3T+sbmEJj9gtngnAkEA
nAdmLlfIG1cqAJGaMfGJZy9nbT6U+Bt89jTWiudVom830oxJwlDRq5AXYugVWjl/
l0JGZQADu3yMJ/rSLFW4QQJAV2cFLL2m5RizxXFlfAB5hm3VWGsF5QzCRuBxC+Ol
TijuczgXpqck/xOHQ8EqCdsR6CiCoXPQzpjyGk+GcHT+Uw==
-----END RSA PRIVATE KEY-----
`
const publicKeyPEM = `
-----BEGIN RSA PUBLIC KEY-----
MIIBCgKCAQEAwsglBy7/slq+/1FllnnFxI7+XFxaenOto9GdAILAAg5OJFXuP0Nh
/WTLz8MN4jYEiDvDNNj+L4xAnCa0WxBsxVc1LseXPzhlqTVjsOYgi0dtKyPl/8tN
VYJb987UYh8TDLJM00dwdZOEkgjQD1rr9Udk+mfGf0+JPV3uguR6a6kCAwEAAQ==
-----END RSA PUBLIC KEY-----
`
