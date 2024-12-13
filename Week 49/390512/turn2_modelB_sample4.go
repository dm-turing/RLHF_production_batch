package main

import (
	"os"
)

type Config struct {
	config string
}

func loadConfig() (*Config, error) {
	if configFile := os.Getenv("CONFIG_FILE"); configFile == "" {
		panic("CONFIG_FILE environment variable is not set")
	}
	// Assume loading config logic
	return &Config{}, nil
}

func main() {
	_, err := loadConfig()
	if err != nil {
		panic("failed to load config: " + err.Error())
	}
	// Proceed with application
}
