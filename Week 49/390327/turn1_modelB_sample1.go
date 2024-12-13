package main

import "fmt"

func processData(data []byte) string {
	cleanedData := data
	result := len(cleanedData)
	return fmt.Sprintf("Result: %d", result)
}

func main() {
	data := "Some sample input"
	fmt.Println(processData([]byte(data)))
}
