package main

import (
	"flag"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"os"
)

// A custom error type that we can use to define specific errors for our application
type customError struct {
	error
}

// Our application function that performs some operation
func doSomething() error {
	// Simulate an error condition
	return customError{errors.New("something went wrong")}
}

func main() {
	// Parse command line flags
	flag.Parse()

	// Perform the main operation
	if err := doSomething(); err != nil {
		// Use the error library to wrap the error and add additional context
		err = errors.Wrap(err, "failed to perform operation")

		// Log the error using the standard log package
		log.Printf("Error: %+v\n", err)

		// Use the error library to extract the underlying error message
		if underlyingErr, ok := errors.Cause(err).(customError); ok {
			// Handle the custom error type specifically
			fmt.Printf("Custom Error: %s\n", underlyingErr.Error())
		} else {
			// Handle other errors generically
			fmt.Printf("Unknown Error: %s\n", err.Error())
		}

		// Exit the application with a non-zero status code
		os.Exit(1)
	}

	// If no error occurred, print a success message
	fmt.Println("Operation successful!")
}