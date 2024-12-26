package main

import (
	"fmt"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func main() {
	// Initialize Viper and set the configuration file location
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	// Read the configuration file
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// Get the configuration value for "application.name"
	appName := viper.GetString("application.name")
	fmt.Println("Application Name:", appName)

	// Watch for configuration changes in real-time
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		appName := viper.GetString("application.name")
		fmt.Println("Updated Application Name:", appName)
	})

	// Simulate updating the configuration file
	go func() {
		for {
			time.Sleep(2 * time.Second)
			viper.Set("application.name", "New App Name")
			viper.WriteConfig()
		}
	}()

	// Simulate running the application for a while
	time.Sleep(10 * time.Second)
}
