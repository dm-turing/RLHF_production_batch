package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// middleware represents a function that will operate on the given input
type middleware func(in string) (string, error)

func logMiddleware(m middleware) middleware {
	return func(in string) (string, error) {
		log.Println("Input:", in)
		return m(in)
	}
}

func uppercaseMiddleware(m middleware) middleware {
	return func(in string) (string, error) {
		return strings.ToUpper(in), nil
	}
}

func containsMiddleware(substring string) middleware {
	return func(in string) (string, error) {
		if !strings.Contains(in, substring) {
			return "", fmt.Errorf("Substring '%s' not found", substring)
		}
		return in, nil
	}
}

func finalHandler(in string) {
	log.Println("Final processed output:", in)
}

func main() {
	log.Println("Welcome to the middleware CLI!")

	// Sequence of operations
	middlewares := []middleware{
		logMiddleware,
		uppercaseMiddleware,
		containsMiddleware("GO"),
	}

	// Scanning user input
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()

		// Composite middleware function sequence
		composedMiddleware := composeMiddleware(middlewares...)

		processedInput, err := composedMiddleware(input)
		if err != nil {
			log.Println("Error processing input:", err)
			continue
		}

		finalHandler(processedInput)
	}

	if err := scanner.Err(); err != nil {
		log.Println("Error scanning input:", err)
	}
}

// composeMiddleware returns a single middleware function that applies all given middleware in order
func composeMiddleware(middlewares ...middleware) middleware {
	return func(in string) (string, error) {
		processedInput := in
		for _, m := range middlewares {
			var err error
			processedInput, err = m(processedInput)
			if err != nil {
				return "", err
			}
		}
		return processedInput, nil
	}
}
