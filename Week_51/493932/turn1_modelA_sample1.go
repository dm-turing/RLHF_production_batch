package main

import (
	"fmt"
)

type Config struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

func main() {
	config := Config{Host: "localhost", Port: 8080}
	jsonString := fmt.Sprintf(`{"host": "%s", "port": %d}`, config.Host, config.Port)
	fmt.Println(jsonString)
}
