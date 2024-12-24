package main
import (
   "fmt"
   "net/http"
   "strings"
   "time"

   "github.com/dgrijalva/jwt-go"
)

const (
   secretKey = "your_secret_key"
)

func generateJWT() (string, error) {
   token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	  "exp": time.Now().Add(time.Hour * 24).Unix(),
	  "sub": "your_subject",
   })
   return token.SignedString([]byte(secretKey))
}

func jwtMiddleware(next http.Handler) http.Handler {
   return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	  // Get the Authorization header
	  authHeader := r.Header.Get("Authorization")
	  if authHeader == "" {
		 http.Error(w, "Unauthorized", http.StatusUnauthorized)
		 return
	  }
	  // Split the header into bearer token
	  tokenString := strings.Split(authHeader, " ")[1]
	  // Parse the JWT token
	  token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		 if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		 }
		 return []byte(secretKey), nil
	  })
	  if err != nil {
		 http.Error(w, "Unauthorized", http.StatusUnauthorized)
		 return
	  }
	  // Validate the token claims
	  if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		 r.Header.Set("sub", claims["sub"].(string))
	  } else {
		 http.Error(w, "Unauthorized", http.StatusUnauthorized)
		 return
	  }
	  // Call the next handler if authorization is successful
	  next.ServeHTTP(w, r)
   })
}

func main() {
   http.HandleFunc("/api/token", func(w http.ResponseWriter, r *http.Request) {