package main

import (
	"fmt"
)

type ActionableError interface {
	error
	action() string
}

type RateLimitError struct {
	Message string
	Action  string
}

func (e *RateLimitError) Error() string {
	return fmt.Sprintf("rate limit error: %s", e.Message)
}

func (e *RateLimitError) action() string {
	return e.Action
}

func checkRateLimit() error {
	// Some rate limit checking logic
	return &RateLimitError{
		Message: "Exceeded hourly rate limit",
		Action:  "Try again in an hour",
	}
}

func main() {
	err := checkRateLimit()
	if aerr, ok := err.(ActionableError); ok {
		fmt.Println("Action required:", aerr.action())
	}
	fmt.Println(err)
}
