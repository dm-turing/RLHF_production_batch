package main

import (
	"fmt"
	"sort"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	sortNumbersCustom(numbers, func(a, b int) bool { return a > b })
	fmt.Println(numbers) // Output: [5 4 3 2 1]
}

func sortNumbersCustom(array []int, compare func(int, int) bool) {
	sort.SliceStable(array, compare)
}
