package main

import (
	"fmt"
	"sort"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	sortNumbers(numbers)
	fmt.Println(numbers) // Output: [1 2 3 4 5]
}

func sortNumbers(array []int) {
	sort.Ints(array)
}
