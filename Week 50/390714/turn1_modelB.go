package main

import (
	"errors"
	"fmt"
	"math"
)

type ErrorType int

const (
	InvalidInputError ErrorType = iota
	InternalServerError
	DatabaseConnectionError
)

func (e ErrorType) String() string {
	switch e {
	case InvalidInputError:
		return "Invalid input error"
	case InternalServerError:
		return "Internal server error"
	case DatabaseConnectionError:
		return "Database connection error"
	default:
		return "Unknown error"
	}
}

func CalculateArea(radius float64) (float64, error) {
	if radius < 0 {
		return 0, errors.New("Radius must be non-negative")
	}

	area := math.Pi * radius * radius

	return area, nil
}

func HandleError(err error) error {
	switch err.(type) {
	case error:
		return err
	default:
		return errors.New("Internal server error")
	}
}

func main() {
	radius := -5.0
	area, err := CalculateArea(radius)
	if err != nil {
		err = HandleError(err)
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Area:", area)
}
