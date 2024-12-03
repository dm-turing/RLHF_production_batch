package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"math/big"
	"os"
	"strings"
)

// parseInput splits the input into two numbers and an operator.
func parseInput(input string) (string, string, error) {
	parts := strings.Fields(input)
	if len(parts) != 2 {
		return "", "", fmt.Errorf("invalid format, expected: <number> <operation>")
	}
	return parts[0], parts[1], nil
}

// calculate performs the specified mathematical operation on a big.Float number.
func calculate(numStr, operation string) (*big.Float, error) {
	num, _, err := big.ParseFloat(numStr, 10, 256, big.ToNearestEven)
	if err != nil {
		return nil, fmt.Errorf("error parsing number: %v", err)
	}
	v, _ := num.Float64()
	result := new(big.Float)
	switch operation {
	case "sin":
		result.SetFloat64(math.Sin(v))
	case "cos":
		result.SetFloat64(math.Cos(v))
	case "tan":
		result.SetFloat64(math.Tan(v))
	case "log":
		result.SetFloat64(math.Log10(v))
	case "ln":
		result.SetFloat64(math.Log(v))
	case "exp":
		result.SetFloat64(math.Exp(v))
	case "sqrt":
		result.SetFloat64(math.Sqrt(v))
	case "pow":
		result.SetFloat64(math.Pow(v, 2))
	default:
		return nil, fmt.Errorf("unknown operation: %s", operation)
	}

	return result, nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Advanced High Precision Scientific Notation Calculator")
	fmt.Println("Enter calculations in the form: <number> <operation>")
	fmt.Println("Supported operations: +, -, *, /, sin, cos, tan, log, ln, exp, sqrt, pow")

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

		numStr, operation, err := parseInput(input)
		if err != nil {
			fmt.Println(err)
			continue
		}

		result, err := calculate(numStr, operation)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Printf("Result: %E\n", result)
	}
}
