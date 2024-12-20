package main

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Database string `envconfig:"DB_URL" required:"true"`
}

func main() {
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		panic(fmt.Errorf("failed to read config from environment: %w", err))
	}

	fmt.Println("Database:", cfg.Database)
}
