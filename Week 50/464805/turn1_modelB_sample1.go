package main

import "fmt"

func main() {
	mySlice := []int{1, 2, 3, 4, 5}
	for _, val := range mySlice {
		fmt.Println(val)
		// Modifying the original slice
		mySlice[0] = 10
	}
}
