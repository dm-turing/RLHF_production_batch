package main

import (
	"fmt"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	filteredAndMapped := make([]int, 0)

	for _, num := range numbers {
		if num%2 == 0 {
			filteredAndMapped = append(filteredAndMapped, num*2)
		}
	}

	fmt.Println(filteredAndMapped)
}
