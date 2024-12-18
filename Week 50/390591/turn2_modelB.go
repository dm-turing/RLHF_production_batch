package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	emailRegex := regexp.MustCompile(`.+@.+\..+`)
	input := strings.Repeat("a", 1000000) + "@example.com"
	if emailRegex.MatchString(input) {
		fmt.Println("Valid email")
	} else {
		fmt.Println("Invalid email")
	}
}
