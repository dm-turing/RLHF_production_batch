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

// parseInput splits the input into a command and arguments.
func parseInput(input string) (string, []string, error) {
	parts := strings.Fields(input)
	if len(parts) < 2 {
		return "", nil, fmt.Errorf("invalid format, expected: <command> <args...>")
	}
	return parts[0], parts[1:], nil
}

// calculate performs the specified operation.
func calculate(command string, args []string) (*big.Float, error) {
	switch command {
	case "add", "sub", "mul", "div", "exp":
		if len(args) != 2 {
			return nil, fmt.Errorf("command %s requires 2 arguments", command)
		}
		num1, _, err := big.ParseFloat(args[0], 10, 256, big.ToNearestEven)
		if err != nil {
			return nil, fmt.Errorf("error parsing first number: %v", err)
		}
		num2, _, err := big.ParseFloat(args[1], 10, 256, big.ToNearestEven)
		if err != nil {
			return nil, fmt.Errorf("error parsing second number: %v", err)
		}

		result := new(big.Float)
		switch command {
		case "add":
			result.Add(num1, num2)
		case "sub":
			result.Sub(num1, num2)
		case "mul":
			result.Mul(num1, num2)
		case "div":
			if num2.Cmp(big.NewFloat(0)) == 0 {
				return nil, fmt.Errorf("division by zero")
			}
			result.Quo(num1, num2)
		case "exp":
			power, _ := num2.Float64()
			f, _ := num1.Float64()
			result.SetFloat64(math.Pow(f, power))
		}
		return result, nil

	case "log", "sin", "cos":
		if len(args) != 1 {
			return nil, fmt.Errorf("command %s requires 1 argument", command)
		}
		num, _, err := big.ParseFloat(args[0], 10, 256, big.ToNearestEven)
		if err != nil {
			return nil, fmt.Errorf("error parsing number: %v", err)
		}

		value, _ := num.Float64()
		result := new(big.Float)

		switch command {
		case "log":
			if value <= 0 {
				return nil, fmt.Errorf("logarithm undefined for non-positive values")
			}
			result.SetFloat64(math.Log(value))
		case "sin":
			result.SetFloat64(math.Sin(value))
		case "cos":
			result.SetFloat64(math.Cos(value))
		}
		return result, nil

	default:
		return nil, fmt.Errorf("unknown command: %s", command)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Advanced Scientific Calculator")
	fmt.Println("Available commands: add, sub, mul, div, exp, log, sin, cos")
	fmt.Println("Enter 'exit' to quit the program.")

	for {
		fmt.Print("Enter command: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Error reading input: %v", err)
		}

		input = strings.TrimSpace(input)

		if strings.ToLower(input) == "exit" {
			fmt.Println("Exiting calculator.")
			break
		}

		command, args, err := parseInput(input)
		if err != nil {
			fmt.Println(err)
			continue
		}

		result, err := calculate(command, args)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Printf("Result: %E\n", result)
	}
}
