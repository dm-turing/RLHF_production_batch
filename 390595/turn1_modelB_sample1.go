package main

import (
	"fmt"
	"os"
)

func main() {
	filePath := "/workspaces/RLHF_production_batch/390595/example.txt"
	dirPath := "/workspaces/RLHF_production_batch/390595/"

	// Check if file exists
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("File does not exist")
		} else {
			fmt.Println("Error checking file:", err)
		}
		return
	}

	if fileInfo.IsDir() {
		fmt.Println("The given path is a directory")
	} else {
		fmt.Println("The given path is a file")
	}

	// Check if directory exists
	dirInfo, err := os.Stat(dirPath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Directory does not exist")
		} else {
			fmt.Println("Error checking directory:", err)
		}
		return
	}

	if dirInfo.IsDir() {
		fmt.Println("The given path is a directory")
	} else {
		fmt.Println("The given path is a file")
	}
}
