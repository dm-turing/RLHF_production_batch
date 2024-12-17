package main

import (
	"fmt"
)

type MyError struct {
	Msg  string
	Code int
}

func (e *MyError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Msg)
}

func someFunction() error {
	return &MyError{Msg: "something failed", Code: 404}
}

func anotherFunction() error {
	if err := someFunction(); err != nil {
		return fmt.Errorf("anotherFunction error: %w", err)
	}
	return nil
}

func main() {
	err := anotherFunction()
	if err != nil {
		fmt.Println(err)
	}
}
