package main

import (
	"fmt"
	"sync"

	"github.com/spf13/viper"
)

// Config represents the configuration struct.
type Config struct {
	Database   string
	ServerPort string
	ApiKey     string
}

var (
	once     sync.Once
	instance *Config
)

// GetConfig returns the configuration instance.
func GetConfig() *Config {
	once.Do(func() {
		viper.SetConfigName("config") // name of config file (without extension)
		viper.AddConfigPath(".")      // search current directory
		// viper.AddConfigPath("/etc/config") // optional: add other search paths
		if err := viper.ReadInConfig(); err != nil {
			panic(fmt.Errorf("failed to read config file: %w", err))
		}

		instance = &Config{
			Database:   viper.GetString("database"),
			ServerPort: viper.GetString("server.port"),
			ApiKey:     viper.GetString("api.key"),
		}
	})

	return instance
}

func main() {
	fmt.Println(GetConfig())
}
