package main

import (
	"errors"
)

type Config struct {
	DatabaseURL string
}

func loadConfig(path string) (Config, error) {
	return Config{}, errors.New("file not found")
}

func init() {
	config, err := loadConfig("config.json")
	if err != nil {
		panic("failed to load configuration: " + err.Error())
	}

	if config.DatabaseURL == "" {
		panic("DatabaseURL must be specified in configuration")
	}
}

func main() {

}
