package main

import (
	"bufio"
	"complex"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseComplex(input string) (complex.Float64, error) {
	parts := strings.Split(input, " ")
	if len(parts) != 2 {
		return 0, fmt.Errorf("invalid format for complex number, expected: <real> <imaginary>")
	}
	realPart, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return 0, fmt.Errorf("error parsing real part: %v", err)
	}
	imaginaryPart, err := strconv.ParseFloat(parts[1], 64)
	if err != nil {
		return 0, fmt.Errorf("error parsing imaginary part: %v", err)
	}
	return complex.Float64(realPart, imaginaryPart), nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Complex Number Calculator")
	fmt.Println("Available operations: add, sub, conjugate")
	fmt.Println("Enter 'exit' to quit the program.")

	for {
		fmt.Print("Enter operation and numbers: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Error reading input: %v", err)
		}

		input = strings.TrimSpace(input)

		if strings.ToLower(input) == "exit" {
			fmt.Println("Exiting calculator.")
			break
		}

		parts := strings.Fields(input)
		if len(parts) < 3 {
			fmt.Println("Invalid input format.")
			continue
		}

		operation := strings.ToLower(parts[0])
		num1, err := parseComplex(parts[1])
		if err != nil {
			fmt.Println(err)
			continue
		}
		num2, err := parseComplex(parts[2])
		if err != nil {
			fmt.Println(err)
			continue
		}

		var result complex.Float64
		switch operation {
		case "add":
			result = num1 + num2
		case "sub":
			result = num1 - num2
		case "conjugate":
			result = complex.Conj(num1)
		default:
			fmt.Println("Unknown operation.")
			continue
		}

		fmt.Printf("Result: %+v\n", result)
	}
}
