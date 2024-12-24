package main

import (
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

func main() {
	// Create the Casbin model and load the policy from a file.
	e, _ := casbin.NewEnforcer("casbin_model.conf", "casbin_policy.csv")

	// Set up the Gin router
	router := gin.Default()

	// Middleware to check for RBAC permissions before allowing requests
	router.Use(CasbinMiddleware(e))

	// Define routes and their associated permissions
	router.GET("/user/:user_id", func(c *gin.Context) {
		userID := c.Param("user_id")
		c.JSON(http.StatusOK, gin.H{"message": "Access granted to user: " + userID})
	})

	router.GET("/admin/dashboard", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Access granted to admin dashboard"})
	})
	router.Run(":8080")
}

// CasbinMiddleware is the middleware function to enforce RBAC permissions.
func CasbinMiddleware(e *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract the user ID from the request context
		userID := c.GetString("user_id")
		// Get the HTTP method and path from the request
		method := c.Request.Method
		path := c.Request.URL.Path

		// Check if the user is allowed to access the requested resource with the given HTTP method
		isAllowed, err := e.Enforce(userID, path, method)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		if !isAllowed {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		// If the user is allowed, proceed with the request
		c.Next()
	}
}
