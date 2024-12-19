package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3, 4, 5}
	for i, x := range s {
		if x > 3 {
			s[i] = 10 // Modifying the slice during iteration
		}
	}
	fmt.Println(s) // Output: [1 2 3 10 10]
}
