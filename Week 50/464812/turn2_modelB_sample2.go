package main

import "fmt"

func main() {
	// Scala-like list comprehension
	numbers := []int{1, 2, 3, 4, 5}
	evenNumbers := []int{}
	for _, num := range numbers {
		if num%2 == 0 {
			evenNumbers = append(evenNumbers, num)
		}
	}
	fmt.Println(evenNumbers)
}
