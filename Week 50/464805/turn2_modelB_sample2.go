package main

import (
	"fmt"
)

func main() {
	s := "Hello, world!"
	for _, c := range s {
		fmt.Printf("%c", c)
		if c == ',' {
			break
		}
	}
	fmt.Println("Reached here!")
}
