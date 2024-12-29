package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

const versionDir = ".filever"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: filever <file_path>")
		return
	}

	filePath := os.Args[1]
	version, err := createVersion(filePath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Created version: %s\n", version)
}

func createVersion(filePath string) (string, error) {
	// Check if the file exists
	_, err := os.Stat(filePath)
	if err != nil {
		return "", err
	}

	// Create the version directory if it doesn't exist
	versionDirPath := filepath.Join(filepath.Dir(filePath), versionDir)
	if _, err := os.Stat(versionDirPath); os.IsNotExist(err) {
		err = os.Mkdir(versionDirPath, 0755)
		if err != nil {
			return "", err
		}
	}

	// Generate the version number based on the current timestamp
	version := time.Now().Format("20060102150405")
	versionPath := filepath.Join(versionDirPath, version)

	// Create a copy of the file with the .version extension
	versionFilePath := filepath.Join(versionPath, filepath.Base(filePath)+".version")
	err = os.MkdirAll(versionPath, 0755)
	if err != nil {
		return "", err
	}

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	err = ioutil.WriteFile(versionFilePath, data, 0644)
	if err != nil {
		return "", err
	}

	return version, nil
}
