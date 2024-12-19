package main

import (
	"fmt"
)

func main() {
	str := "Hello, 世界"
	for index, rune := range str {
		fmt.Printf("Index: %d, Rune: %c\n", index, rune)
	}
}
