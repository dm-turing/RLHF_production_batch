package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func validateEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

func validateUSPhoneNumber(phone string) bool {
	phoneRegex := regexp.MustCompile(`^\(?[0-9]{3}\)?[\s.-]?[0-9]{3}[\s.-]?[0-9]{4}$`)
	return phoneRegex.MatchString(phone)
}

func validateNumberRange(num string, min, max int) bool {
	numInt, err := strconv.Atoi(num)
	if err != nil {
		return false
	}
	return numInt >= min && numInt <= max
}

func validateFloatRange(num string, min, max float64) bool {
	numFloat, err := strconv.ParseFloat(num, 64)
	if err != nil {
		return false
	}
	return numFloat >= min && numFloat <= max
}

func validateDateFormat(date string) bool {
	dateRegex := regexp.MustCompile(`^(\d{4}-\d{2}-\d{2})$`)
	return dateRegex.MatchString(date)
}

func main() {
	emails := []string{"example@example.com", "InvalidEmail", "Example.email@example.com"}
	phoneNumbers := []string{"(555) 123-4567", "555 123 4567", "5551234567", "InvalidPhone"}
	numbers := []string{"10", "20", "30", "31", "32"}
	floats := []string{"10.1", "20.2", "30.3", "31.4", "32.5"}
	dates := []string{"2023-01-01", "2023-01-32", "2023-02-29", "2023-04-31", "InvalidDate"}

	fmt.Println("Email Validation:")
	for _, email := range emails {
		if validateEmail(email) {
			fmt.Printf("Email %s is valid\n", email)
		} else {
			fmt.Printf("Email %s is invalid\n", email)
		}
	}

	fmt.Println("\nPhone Number Validation:")
	for _, phone := range phoneNumbers {
		if validateUSPhoneNumber(phone) {
			fmt.Printf("Phone %s is valid\n", phone)
		} else {
			fmt.Printf("Phone %s is invalid\n", phone)
		}
	}

	fmt.Println("\nNumber Range Validation:")
	for _, number := range numbers {
		if validateNumberRange(number, 10, 30) {
			fmt.Printf("Number %s is valid\n", number)
		} else {
			fmt.Printf("Number %s is invalid\n", number)
		}
	}

	fmt.Println("\nFloat Range Validation:")
	for _, number := range floats {
		if validateFloatRange(number, 10.0, 30.0) {
			fmt.Printf("Float %s is valid\n", number)
		} else {
			fmt.Printf("Float %s is invalid\n", number)
		}
	}

	fmt.Println("\nDate Format Validation:")
	for _, date := range dates {
		if validateDateFormat(date) {
			fmt.Printf("Date %s is valid\n", date)
		} else {
			fmt.Printf("Date %s is invalid\n", date)
		}
	}
}
