// Go code (comparison with Java)
package main

import "fmt"

func main() {
	name := "Bob"
	age := 35
	// Java equivalent: String greeting = String.format("Hello, %s! You are %d years old.", name, age);
	greeting := fmt.Sprintf("Hello, %s! You are %d years old.", name, age)
	fmt.Println(greeting)
}
