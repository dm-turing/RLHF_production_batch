package main

import (
	"fmt"
	"os"
)

// Configuration represents the structure of the configuration.
type Configuration struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
}

func main() {
	// Sample configuration data.
	config := Configuration{
		Host:     "localhost",
		Port:     5432,
		Username: "dbuser",
		Password: "secret",
		Database: "mydb",
	}

	// Use fmt.Sprintf to dynamically format the configuration string.
	configString := fmt.Sprintf(
		`[database]
host = "%s"
port = %d
username = "%s"
password = "%s"
database = "%s"
`,
		config.Host,
		config.Port,
		config.Username,
		config.Password,
		config.Database,
	)

	// Write the configuration string to a file.
	configFile, err := os.Create("config.ini")
	if err != nil {
		panic(fmt.Errorf("could not create file: %v", err))
	}
	defer configFile.Close()

	_, err = configFile.Write([]byte(configString))
	if err != nil {
		panic(fmt.Errorf("could not write to file: %v", err))
	}

	fmt.Println("Configuration file generated successfully.")
}
