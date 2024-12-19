package main

import (
	"fmt"
)

func main() {
	userData := map[string]int{"Alice": 25, "Bob": 30}
	for key, value := range userData {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}
}
