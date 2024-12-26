package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	DBHost    string `env:"DB_HOST" default:"localhost"`
	DBPort    int    `env:"DB_PORT" default:5432`
	LogLevel  string `env:"LOG_LEVEL" default:"info"`
	DebugMode bool   `env:"DEBUG_MODE" default:false`
}

func loadConfigFromFile(configFile string) (Config, error) {
	// Assuming configuration is in key=value format
	var config Config

	file, err := os.Open(configFile)
	if err != nil {
		return config, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") {
			continue // Skip comments
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])

			switch key {
			case "DB_HOST":
				config.DBHost = value
			case "DB_PORT":
				port, err := strconv.Atoi(value)
				if err == nil {
					config.DBPort = port
				}
			case "LOG_LEVEL":
				config.LogLevel = value
			case "DEBUG_MODE":
				config.DebugMode, _ = strconv.ParseBool(value)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return config, err
	}

	return config, nil
}

func loadConfig() (Config, error) {
	configFile := os.Getenv("CONFIG_FILE")
	if configFile == "" {
		configFile = "config.conf"
	}

	return loadConfigFromFile(configFile)
}

func main() {
	config, err := loadConfig()
	if err != nil {
		fmt.Println("Error loading config:", err)
		os.Exit(1)
	}

	// Print initial configuration
	fmt.Println("Initial Config:", config)

	// Reload configuration every 5 seconds
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		newConfig, err := loadConfig()
		if err != nil {
			fmt.Println("Error reloading config:", err)
			continue
		}

		if config.DBHost != newConfig.DBHost {
			fmt.Println("DB_HOST updated from", config.DBHost, "to", newConfig.DBHost)
		}

		config = newConfig
	}
}
