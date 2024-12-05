package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Define a flag for the maximum value
	maxValue := flag.Int("max", 100, "The maximum value for the random number")
	flag.Parse()

	// Parse additional arguments if given
	if len(os.Args) > 1 {
		arg := os.Args[1]
		if maxVal, err := strconv.Atoi(arg); err == nil {
			*maxValue = maxVal
		} else {
			fmt.Println("Invalid number provided, using default maximum value of 100")
		}
	}

	// Generate the random number
	randomNumber := rand.Intn(*maxValue + 1)
	fmt.Printf("Random number between 0 and %d: %d\n", *maxValue, randomNumber)
}
