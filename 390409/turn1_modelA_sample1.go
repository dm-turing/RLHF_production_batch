package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func generateRandomNumbers(min, max int, count int) {
	for i := 0; i < count; i++ {
		fmt.Printf("%d ", rand.Intn(max-min+1)+min)
	}
	fmt.Println()
}

func main() {
	rand.Seed(int64(os.Getpid()))

	if len(os.Args) != 4 {
		fmt.Println("Usage: go run random_number.go <min> <max> <count>")
		return
	}

	min, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid input for min")
		return
	}

	max, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Invalid input for max")
		return
	}

	count, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Println("Invalid input for count")
		return
	}

	if min > max {
		fmt.Println("Minimum must be less than or equal to maximum")
		return
	}

	generateRandomNumbers(min, max, count)
}
