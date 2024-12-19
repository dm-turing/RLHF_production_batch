package main

import (
	"fmt"
)

func main() {
	capitals := map[string]string{"France": "Paris", "Italy": "Rome", "Germany": "Berlin"}
	for key, value := range capitals {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}
}
