package main

import (
	"fmt"
	"strings"
)

// Named function filter
func filterStrings(Strings []string, f func(string) bool) []string {
	var filtered []string
	for _, s := range Strings {
		if f(s) {
			filtered = append(filtered, s)
		}
	}
	return filtered
}

func main() {
	Strings := []string{"apple", "banana", "cherry", "orange"}

	// Using a named function
	isEvenLength := func(s string) bool {
		return len(s)%2 == 0
	}
	evenLengthStrings := filterStrings(Strings, isEvenLength)
	fmt.Println("Even length strings:", evenLengthStrings) // Output: Even length strings: [banana cherry]

	// Using an anonymous function
	containsVowel := func(s string) bool {
		for _, c := range s {

			if strings.Contains("aeiouAEIOU", string(c)) {
				return true
			}
		}
		return false
	}
	stringsWithVowels := filterStrings(Strings, containsVowel)
	fmt.Println("Strings with vowels:", stringsWithVowels) // Output: Strings with vowels: [apple banana orange]
}
