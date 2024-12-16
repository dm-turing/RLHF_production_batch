package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	user := User{ID: 1, Name: "John"}
	jsonBytes, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonBytes)) // Output: {"id":1,"name":"John"}
}
