package main

import (
	"fmt"
)

type ValidationError struct {
	Field   string
	Message string
}

func validateUserInput(name, email string) (errors []ValidationError) {
	if name == "" {
		errors = append(errors, ValidationError{"Name", "is required"})
	}

	if email == "" {
		errors = append(errors, ValidationError{"Email", "is required"})
	}

	// Simulate an email validation check
	if email != "valid@example.com" {
		errors = append(errors, ValidationError{"Email", "is invalid"})
	}

	return
}

func handleValidationErrors(errors []ValidationError) {
	if len(errors) == 0 {
		fmt.Println("Input is valid")
	} else {
		// Decision tree logic
		for _, error := range errors {
			switch error.Field {
			case "Name":
				switch error.Message {
				case "is required":
					fmt.Println("Please enter a name.")
				default:
					fmt.Println("Unknown name error:", error.Message)
				}
			case "Email":
				switch error.Message {
				case "is required":
					fmt.Println("Please enter an email.")
				case "is invalid":
					fmt.Println("Invalid email format. Please try again.")
				default:
					fmt.Println("Unknown email error:", error.Message)
				}
			default:
				fmt.Println("Unknown field error:", error.Field)
			}
		}
	}
}

func main() {
	name := ""
	email := "invalid@example.com"

	errors := validateUserInput(name, email)
	handleValidationErrors(errors)
}
