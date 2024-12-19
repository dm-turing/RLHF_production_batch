package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3, 4, 5}
	var result []int
	for _, x := range s {
		if x > 3 {
			result = append(result, x*2)
		}
	}
	fmt.Println(result) // Output: [8 10]
}
