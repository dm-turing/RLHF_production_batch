package main

import (
	"fmt"
)

func main() {
	name := "Alice"
	age := 30
	language := "Go"
	programmingYears := 5
	formattedString := fmt.Sprintf("Hello, my name is %s. I am %d years old and have been programming in %s for %d years.", name, age, language, programmingYears)
	fmt.Println(formattedString)
}
