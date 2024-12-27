package main

import (
	"errors"
	"fmt"
	"runtime"
)

type LoggingError struct {
	Message  string
	Function string
	Line     int
}

func (e *LoggingError) Error() string {
	return fmt.Sprintf("%s at %s:%d", e.Message, e.Function, e.Line)
}

func someFunction() error {
	err := errors.New("an error occurred")
	return &LoggingError{
		Message:  err.Error(),
		Function: "someFunction",
		Line: (func() int {
			_, _, line, _ := runtime.Caller(1)
			return line
		})(),
	}
}

func main() {
	fmt.Println(someFunction().Error())
}
