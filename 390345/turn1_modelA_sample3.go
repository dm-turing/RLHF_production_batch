package main

import (
	"fmt"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	evenSum := 0

	// Create a channel to send even numbers
	evenNumbers := make(chan int)

	// Goroutine to filter even numbers
	go func() {
		for _, num := range numbers {
			if num%2 == 0 {
				evenNumbers <- num
			}
		}
		close(evenNumbers)
	}()

	// Goroutine to sum even numbers
	go func() {
		for evenNum := range evenNumbers {
			evenSum += evenNum
		}
		fmt.Println("Sum of even numbers:", evenSum)
	}()

	// Wait for both goroutines to complete
	<-evenNumbers
}
