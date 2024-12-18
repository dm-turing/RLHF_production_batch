package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// CORS Middleware Configuration
	r.Use(cors.New(cors.Config{
		// Allowed origins - specify the domains that can make requests
		AllowOrigins:     []string{"https://example.com", "https://anotherdomain.com"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		MaxAge:           86400, // Allow the preflight request cache to live for 24 hours
	}))

	// Example API endpoint
	r.GET("/api/data", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello from the API!"})
	})

	// Start the server
	r.Run(":8080")
}
