package main

import (
	"fmt"
)

func main() {
	var slice []int
	for i := 0; i < 1000000; i++ {
		slice = append(slice, i)
	}

	// Setting the length to 0 does not free the underlying array
	slice = slice[:0]
	fmt.Println("Slice after setting length to 0:", cap(slice))
}
