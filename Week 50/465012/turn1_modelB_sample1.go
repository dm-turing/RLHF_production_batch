package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

// Configuration struct holds the configuration settings for the web server.
type Configuration struct {
	Host string
	Port int
	SSL  bool
}

// generateConfig generates a configuration file content based on the given Configuration struct.
func generateConfig(config Configuration) string {
	configTemplate := `
Listen %s:%d
SSL %t
`
	return fmt.Sprintf(configTemplate, config.Host, config.Port, config.SSL)
}

func main() {
	// Create a Configuration struct with your desired settings.
	config := Configuration{
		Host: "localhost",
		Port: 8080,
		SSL:  false,
	}

	// Generate the configuration file content using the generateConfig function.
	configContent := generateConfig(config)

	// Write the configuration content to a file named "webserver.conf"
	err := ioutil.WriteFile("webserver.conf", []byte(configContent), 0644)
	if err != nil {
		log.Fatalf("Error writing configuration file: %v", err)
	}

	fmt.Println("Configuration file generated successfully!")
}
