package main

import (
	"fmt"
	"strings"
)

func uppercase(s string) string {
	return strings.ToUpper(s)
}

func trim(s string) string {
	return strings.TrimSpace(s)
}

func reverse(s string) string {
	runes := []rune(s)
	for i := 0; i < len(runes)/2; i++ {
		j := len(runes) - i - 1
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func compose(f, g func(string) string) func(string) string {
	return func(s string) string {
		return f(g(s))
	}
}

func main() {
	transform := compose(uppercase, compose(reverse, trim))
	result := transform("  hello world  ")
	fmt.Println(result) // Output: DLROW OLLEH
}
