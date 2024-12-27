package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"

	"gopkg.in/yaml.v2"
)

// Configuration represents the data structure of our configuration
type Configuration struct {
	Database struct {
		Host     string `json:"host" xml:"host" yaml:"host"`
		Port     int    `json:"port" xml:"port" yaml:"port"`
		User     string `json:"user" xml:"user" yaml:"user"`
		Password string `json:"password" xml:"password" yaml:"password"`
		DBName   string `json:"dbname" xml:"dbname" yaml:"dbname"`
	} `json:"database" xml:"database" yaml:"database"`
	Application struct {
		Name    string `json:"name" xml:"name" yaml:"name"`
		Port    int    `json:"port" xml:"port" yaml:"port"`
		Env     string `json:"env" xml:"env" yaml:"env"`
		Timeout int    `json:"timeout" xml:"timeout" yaml:"timeout"`
	} `json:"application" xml:"application" yaml:"application"`
}

func main() {
	// Sample configuration data
	config := Configuration{
		Database: struct {
			Host     string `json:"host" xml:"host" yaml:"host"`
			Port     int    `json:"port" xml:"port" yaml:"port"`
			User     string `json:"user" xml:"user" yaml:"user"`
			Password string `json:"password" xml:"password" yaml:"password"`
			DBName   string `json:"dbname" xml:"dbname" yaml:"dbname"`
		}{
			Host:     "localhost",
			Port:     5432,
			User:     "postgres",
			Password: "secret",
			DBName:   "mydb",
		},
		Application: struct {
			Name    string `json:"name" xml:"name" yaml:"name"`
			Port    int    `json:"port" xml:"port" yaml:"port"`
			Env     string `json:"env" xml:"env" yaml:"env"`
			Timeout int    `json:"timeout" xml:"timeout" yaml:"timeout"`
		}{
			Name:    "MyApp",
			Port:    8080,
			Env:     "development",
			Timeout: 30,
		},
	}
	// JSON Format
	jsonBytes, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		fmt.Println("Error generating JSON configuration:", err)
	} else {
		fmt.Println("Generated JSON Configuration:")
		fmt.Println(string(jsonBytes))
	}
	// XML Format
	xmlBytes, err := xml.MarshalIndent(config, "", "  ")
	if err != nil {
		fmt.Println("Error generating XML configuration:", err)
	} else {
		fmt.Println("\nGenerated XML Configuration:")
		fmt.Println(string(xmlBytes))
	}

	// YAML Format
	yamlBytes, err := yaml.Marshal(config)
	if err != nil {
		fmt.Println("Error generating YAML configuration:", err)
	} else {
		fmt.Println("\nGenerated YAML Configuration:")
		fmt.Println(string(yamlBytes))
	}
}
