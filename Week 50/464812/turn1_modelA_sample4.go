package main

import "fmt"

func main() {
	greeting := "Hello, World!"
	for _, char := range greeting {
		fmt.Print(char)
	}
}
