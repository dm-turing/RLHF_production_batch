package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Set up CORS middleware
	// Ensure that only allowed origins are allowed to make requests
	// AllowList adds domains to the list of allowed origins
	router.Use(CORSMiddleware([]string{"https://example.com", "https://subdomain.example.com"}))

	// Handle your API endpoints here
	router.GET("/api/products", func(c *gin.Context) {
		// Handle API request
		// ...
	})

	router.Run(":8080")
}

// CORSMiddleware sets up CORS headers for the given origins.
func CORSMiddleware(allowList []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the origin header
		origin := c.Request.Header.Get("Origin")

		// Check if the origin is in the allow list
		// If it is, allow the request
		for _, allowedOrigin := range allowList {
			if allowedOrigin == origin {
				c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
				break
			}
		}

		// Set the other CORS headers
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		// Handle preflight requests
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
