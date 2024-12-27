package main

import (
	"fmt"
)

type Config struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

func main() {
	config := Config{Host: "localhost", Port: 8080}
	yamlString := fmt.Sprintf(`host: %s\nport: %d\n`, config.Host, config.Port)
	fmt.Println(yamlString)
}
