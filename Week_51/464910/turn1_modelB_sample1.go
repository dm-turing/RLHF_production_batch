package main

import (
	"fmt"
)

type Callback func(int)

func processData(data []int, callback Callback) {
	for _, value := range data {
		callback(value)
	}
}

func main() {
	data := []int{1, 2, 3, 4, 5}

	processData(data, func(value int) {
		fmt.Println("Callback value:", value)
	})
}
