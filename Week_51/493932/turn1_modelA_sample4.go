package main

import (
	"encoding/json"
	"fmt"
)

type Config struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

func main() {
	config := Config{Host: "localhost", Port: 8080}
	jsonBytes, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(jsonBytes))
}
