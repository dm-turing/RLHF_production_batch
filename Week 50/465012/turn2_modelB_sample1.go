package main

import (
	"fmt"
	"log"

	"github.com/pelletier/go-toml"
)

type Configuration struct {
	Server struct {
		Host string `toml:"host"`
		Port int    `toml:"port"`
		SSL  bool   `toml:"ssl"`
	}
	Database struct {
		Host     string `toml:"host"`
		Port     int    `toml:"port"`
		User     string `toml:"user"`
		Password string `toml:"password"`
		Name     string `toml:"name"`
	}
}

func main() {
	var config Configuration
	c, _ := toml.LoadFile("config.toml")
	d, _ := toml.Marshal(c)
	if err := toml.Unmarshal(d, &config); err != nil {
		log.Fatalf("Error decoding configuration file: %v", err)
	}

	// Now you can access the configuration fields using the Configuration struct
	fmt.Println(config.Server.Host)
	fmt.Println(config.Database.Host)
}
