package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "myapp",
	Short: "A brief description of your app",
	Run: func(cmd *cobra.Command, args []string) {
		host := viper.GetString("host")
		port := viper.GetInt("port")
		fmt.Printf("Server host: %s, port: %d\n", host, port)
	},
}

func main() {
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config/")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s\n", err))
	}

	rootCmd.Execute()
}
