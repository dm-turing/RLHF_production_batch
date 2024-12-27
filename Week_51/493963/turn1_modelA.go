package main

import "fmt"

type ValidationError struct {
	Field  string
	Reason string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation error: %s on field %s", e.Reason, e.Field)
}

type DatabaseError struct {
	Cause   error
	Query   string
	ErrCode int
}

func (e *DatabaseError) Error() string {
	return fmt.Sprintf("database error: %s (%d) while executing: %s", e.Cause, e.ErrCode, e.Query)
}

func someFunction() error {
	return fmt.Errorf("something went wrong: %w", someOtherFunction())
}

func someOtherFunction() error {
	return fmt.Errorf("%s", "some other error")
}

func main() {
	v := ValidationError{
		Field:  "Error1",
		Reason: "Because error 1 happened",
	}
	fmt.Println(v.Error())
	d := DatabaseError{
		Cause:   fmt.Errorf("%s", "database error"),
		Query:   "during write operation",
		ErrCode: 101,
	}
	fmt.Println(d.Error())
	fmt.Println(someFunction())
}
