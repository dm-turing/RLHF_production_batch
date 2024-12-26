package main

import (
	"fmt"
	"strings"
)

func main() {
	names := []string{"Alice", "Bob", "Charlie"}
	// Using a named function for filtering
	isEvenLength := func(s string) bool {
		return len(s)%2 == 0
	}
	evenLengthNames := filterNames(names, isEvenLength)
	fmt.Println("Even length names:", strings.Join(evenLengthNames, ", "))

	// Using an anonymous function for filtering
	oddLengthNames := filterNames(names, func(s string) bool {
		return len(s)%2 != 0
	})
	fmt.Println("Odd length names:", strings.Join(oddLengthNames, ", "))
}
func filterNames(names []string, predicate func(string) bool) []string {
	var filteredNames []string
	for _, name := range names {
		if predicate(name) {
			filteredNames = append(filteredNames, name)
		}
	}
	return filteredNames
}
