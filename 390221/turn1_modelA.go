package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter an expression (use e for scientific notation, e.g., 3.14e10):")

	for scanner.Scan() {
		input := scanner.Text()
		parts := strings.Split(input, " ")

		if len(parts) != 2 {
			fmt.Println("Invalid input. Use the format: 'operation num1 num2'")
			continue
		}

		operation := parts[0]
		num1Str := strings.TrimSpace(parts[1])

		// Convert input to a big.Float
		num1, ok := new(big.Float).SetString(num1Str)
		if !ok {
			fmt.Println("Invalid number format")
			continue
		}

		// Get the second number from the user
		fmt.Print("Enter the second number: ")
		scanner.Scan()
		num2Str := scanner.Text()

		num2, ok := new(big.Float).SetString(num2Str)
		if !ok {
			fmt.Println("Invalid number format")
			continue
		}

		// Perform the operation
		var result *big.Float
		switch operation {
		case "+":
			result = num1.Add(num1, num2)
		case "-":
			result = num1.Sub(num1, num2)
		case "*":
			result = num1.Mul(num1, num2)
		case "/":
			if num2.Cmp(big.NewFloat(0)) == 0 {
				fmt.Println("Error: Division by zero")
				continue
			}
			result = num1.Quo(num1, num2)
		default:
			fmt.Println("Invalid operation")
			continue
		}

		// Print the result
		fmt.Printf("Result: %s\n", result.String())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "scanning error:", err)
	}
}
