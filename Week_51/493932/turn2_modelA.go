package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-playground/validator/v10"
)

// Define the data structure
type ComplexConfig struct {
	Host     string         `json:"host" validate:"required,nonzero"`
	Port     int            `json:"port" validate:"required,min=1"`
	Services []Service      `json:"services" validate:"required"`
	Logging  *LoggingConfig `json:"logging"`
}

type Service struct {
	Name     string `json:"name" validate:"required"`
	Enabled  bool   `json:"enabled"`
	Endpoint string `json:"endpoint"`
}

type LoggingConfig struct {
	Level string `json:"level" validate:"oneof=debug info warning error"`
	File  string `json:"file"`
}

// Validate the data
func (c *ComplexConfig) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}

// Format the data into a string using fmt.Sprintf
func (c *ComplexConfig) FormatJSON() string {
	bytes, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		log.Fatalf("Error formatting JSON: %v", err)
	}
	return string(bytes)
}

// Write the formatted string to a file
func WriteConfigToFile(filename string, configString string) error {
	err := fmt.Printf("%s", configString)
	if err != nil {
		log.Fatalf("Error writing to file: %v", err)
	}
	return err
}

func main() {
	// Create a complex config instance
	config := ComplexConfig{
		Host: "localhost",
		Port: 8080,
		Services: []Service{
			{Name: "service1", Enabled: true, Endpoint: "/api/v1/service1"},
			{Name: "service2", Enabled: false, Endpoint: "/api/v1/service2"},
		},
		Logging: &LoggingConfig{Level: "info", File: "/var/log/app.log"},
	}

	// Validate the config
	if err := config.Validate(); err != nil {
		log.Fatalf("Invalid config: %v", err)
	}

	// Format the config as JSON
	jsonConfigString := config.FormatJSON()

	// Write the formatted JSON to a file
	WriteConfigToFile("config.json", jsonConfigString)
	log.Println("Configuration file generated successfully.")
}
