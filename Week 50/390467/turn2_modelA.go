package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// CustomError defines a custom error type for API errors.
type CustomError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Error implements the error interface.
func (e *CustomError) Error() string {
	return e.Message
}

// NewCustomError creates a new CustomError.
func NewCustomError(code int, message string) *CustomError {
	return &CustomError{
		Code:    code,
		Message: message,
	}
}

// Global logger
var logger = logrus.New()

// ErrorHandlingMiddleware is a middleware to handle errors centrally
func ErrorHandlingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next() // Process the request

		// Check if there were any errors
		if len(c.Errors) > 0 {
			var customErr *CustomError
			for _, err := range c.Errors {
				if e, ok := err.Err.(*CustomError); ok {
					customErr = e
					break
				}
			}
			if customErr != nil {
				logger.WithFields(logrus.Fields{
					"status": customErr.Code,
					"error":  customErr.Message,
				}).Error("API Error")

				c.JSON(customErr.Code, gin.H{
					"error": customErr.Message,
				})
			} else {
				logger.Error("Internal Server Error")
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Internal Server Error",
				})
			}
		}
	}
}

// ExampleHandler is a sample handler that generates an error
func ExampleHandler(c *gin.Context) {
	// Simulating a validation error
	if c.Query("name") == "" {
		c.Error(NewCustomError(http.StatusBadRequest, "Name is required"))
		return
	}

	// Normal response
	c.JSON(http.StatusOK, gin.H{
		"message": "Success!",
	})
}

// Another example handler with a database error simulation
func DatabaseHandler(c *gin.Context) {
	// Simulating a database error
	c.Error(NewCustomError(http.StatusInternalServerError, "Database connection failed"))
}

func main() {
	// Set up Gin router
	router := gin.Default()

	// Use the error handling middleware
	router.Use(ErrorHandlingMiddleware())

	// Define routes
	router.GET("/example", ExampleHandler)
	router.GET("/database", DatabaseHandler)

	// Start the server
	if err := router.Run(":8080"); err != nil {
		logger.Fatalf("Failed to start server: %v", err)
	}
}
