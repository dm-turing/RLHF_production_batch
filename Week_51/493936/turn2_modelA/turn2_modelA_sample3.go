package main

import (
	"fmt"
	"os"
)

type ErrorType interface {
	error
	Type() string
}

type FatalError struct {
	message string
}

func (e *FatalError) Error() string {
	return e.message
}

func (e *FatalError) Type() string {
	return "fatal"
}

func main() {
	f := FatalError{message: "application crash"}
	if f.Type() == "fatal" {
		fmt.Println("Fatal error encountered!")
		os.Exit(1)
	}
}
