package main

import "fmt"

type ComplexConfig struct {
	Server   ServerConfig   `json:"server"`
	Database DatabaseConfig `json:"database"`
}

type ServerConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
	SSL  bool   `json:"ssl"`
}

type DatabaseConfig struct {
	Type     string `json:"type"`
	User     string `json:"user"`
	Password string `json:"password"`
}

func generateComplexConfig() (ComplexConfig, error) {
	// Create the configuration data using structs
	config := ComplexConfig{
		Server: ServerConfig{
			Host: "localhost",
			Port: 8080,
			SSL:  false,
		},
		Database: DatabaseConfig{
			Type:     "postgres",
			User:     "root",
			Password: "secret",
		},
	}
	return config, nil
}

func main() {
	fmt.Println(generateComplexConfig())
}
