package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
)

// Middleware function to validate CLI input
func Middleware(next func() error) error {
	flag.Parse()
	if *operation == "" {
		log.Println("Operation is required.")
		flag.PrintDefaults()
		return nil
	}
	return next()
}

var (
	operation = flag.String("op", "", "Operation to perform (add, sub, mul, div)")
	num1      = flag.Float64("num1", 0, "First number")
)

func add() error {
	a, _ := strconv.ParseFloat(flag.Arg(0), 32)
	b, _ := strconv.ParseFloat(flag.Arg(1), 32)
	fmt.Printf("Result: %.2f\n", a+b)
	return nil
}

func sub() error {
	a, _ := strconv.ParseFloat(flag.Arg(0), 32)
	b, _ := strconv.ParseFloat(flag.Arg(1), 32)
	fmt.Printf("Result: %.2f\n", a-b)
	return nil
}

func mul() error {
	a, _ := strconv.ParseFloat(flag.Arg(0), 32)
	b, _ := strconv.ParseFloat(flag.Arg(1), 32)
	fmt.Printf("Result: %.2f\n", a*b)
	return nil
}

func div() error {
	if flag.Arg(1) == "0" {
		return fmt.Errorf("cannot divide by zero")
	}
	a, _ := strconv.ParseFloat(flag.Arg(0), 32)
	b, _ := strconv.ParseFloat(flag.Arg(1), 32)
	fmt.Printf("Result: %.2f\n", a/b)
	return nil
}

func main() {
	if err := Middleware(func() error {
		switch *operation {
		case "add":
			return add()
		case "sub":
			return sub()
		case "mul":
			return mul()
		case "div":
			return div()
		default:
			return fmt.Errorf("unknown operation: %s", *operation)
		}
	}); err != nil {
		log.Fatal(err)
	}
}
