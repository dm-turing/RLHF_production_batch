package main

import "encoding/json"

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email,omitempty"`
}

func main() {
	var people []Person
	jsonData := `[{"name":"Alice","age":30},{"name":"Bob","age":25}]`
	json.Unmarshal([]byte(jsonData), &people)
}
