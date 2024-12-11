package main

import (
	"errors"
	"fmt"
	"runtime"
	"time"
)

type CustomError struct {
	Message    string
	Code       int
	Cause      error
	StackTrace []byte
	Timestamp  time.Time
	Context    map[string]interface{}
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Message)
}

func NewCustomError(message string, code int, cause error, context map[string]interface{}) error {
	// Capture stack trace
	buf := make([]byte, 4096)
	buf = buf[:runtime.Stack(buf, false)]
	e := &CustomError{
		Message:    message,
		Code:       code,
		Cause:      cause,
		StackTrace: buf,
		Timestamp:  time.Now(),
		Context:    context,
	}
	return e
}

func main() {
	context := map[string]interface{}{
		"user_id":    "123",
		"request_id": "abc",
	}

	err := NewCustomError("Invalid input", 400, errors.New("value cannot be nil"), context)

	if customErr, ok := err.(*CustomError); ok {
		fmt.Println("Custom Error Code:", customErr.Code)
		fmt.Println("Error Timestamp:", customErr.Timestamp)
		fmt.Println("Context:", customErr.Context)
		fmt.Println("Stack Trace:", string(customErr.StackTrace))
	} else {
		fmt.Println("Unknown error")
	}
}
