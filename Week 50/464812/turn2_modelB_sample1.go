package main

import "fmt"

func main() {
	// Python-like for-in loop over a slice
	numbers := []int{1, 2, 3, 4, 5}
	for _, num := range numbers {
		fmt.Println(num)
	}

	// Python-like for-in loop over a string
	greeting := "Hello, World!"
	for _, char := range greeting {
		fmt.Print(string(char))
	}
}
