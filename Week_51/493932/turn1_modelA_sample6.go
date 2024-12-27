package main

import (
	"fmt"
	"log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

func main() {
	config := Config{Host: "localhost", Port: 8080}
	yamlBytes, err := yaml.Marshal(config)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(yamlBytes))
}
