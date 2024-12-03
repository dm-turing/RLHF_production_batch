package main

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"
)

// parseInput splits the input into two numbers and an operator.
func parseInput(input string) (string, string, string, error) {
	parts := strings.Fields(input)
	if len(parts) != 3 {
		return "", "", "", fmt.Errorf("invalid format, expected: <number> <operator> <number>")
	}
	return parts[0], parts[1], parts[2], nil
}

// calculate performs the arithmetic operation on two big.Float numbers.
func calculate(num1Str, operator, num2Str string) (*big.Float, error) {
	num1, _, err := big.ParseFloat(num1Str, 10, 256, big.ToNearestEven)
	if err != nil {
		return nil, fmt.Errorf("error parsing first number: %v", err)
	}
	num2, _, err := big.ParseFloat(num2Str, 10, 256, big.ToNearestEven)
	if err != nil {
		return nil, fmt.Errorf("error parsing second number: %v", err)
	}

	result := new(big.Float)
	switch operator {
	case "+":
		result.Add(num1, num2)
	case "-":
		result.Sub(num1, num2)
	case "*":
		result.Mul(num1, num2)
	case "/":
		if num2.Cmp(big.NewFloat(0)) == 0 {
			return nil, fmt.Errorf("division by zero")
		}
		result.Quo(num1, num2)
	default:
		return nil, fmt.Errorf("unknown operator: %s", operator)
	}

	return result, nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("High Precision Scientific Notation Calculator")
	fmt.Println("Enter calculations in the form: <number1> <operator> <number2>")
	fmt.Println("Supported operators: +, -, *, /")

	for {
		fmt.Print("Enter calculation: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Error reading input: %v", err)
		}

		input = strings.TrimSpace(input)

		if strings.ToLower(input) == "exit" {
			fmt.Println("Exiting calculator.")
			break
		}

		num1Str, operator, num2Str, err := parseInput(input)
		if err != nil {
			fmt.Println(err)
			continue
		}

		result, err := calculate(num1Str, operator, num2Str)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Printf("Result: %E\n", result)
	}
}
