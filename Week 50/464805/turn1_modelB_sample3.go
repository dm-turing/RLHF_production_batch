package main

import "fmt"

func main() {
	mySlice := make([]int, 10)
	for i := range mySlice {
		mySlice[i] = 10
	}
	fmt.Println(mySlice)
}
