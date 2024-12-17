package main

import (
	"errors"
	"fmt"
)

type MyError struct{}

func (e *MyError) Error() string {
	return "my special error happened"
}

func mightCauseError() error {
	return &MyError{}
}

func doSomething() error {
	if err := mightCauseError(); err != nil {
		return fmt.Errorf("doSomething encountered an error: %w", err)
	}
	return nil
}

func main() {
	err := doSomething()
	if err != nil {
		fmt.Println("Error:", err)

		var myErr *MyError
		if errors.As(err, &myErr) {
			fmt.Println("Caught MyError")
		}
	}
}
