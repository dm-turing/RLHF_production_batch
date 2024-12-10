package main

import (
	"fmt"
	"os"
)

func main() {
	filePath := "/workspaces/RLHF_production_batch/390595/example.txt"

	// Check if file exists and has read permissions
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("File does not exist")
		} else {
			fmt.Println("Error checking file:", err)
		}
		return
	}

	if fileInfo.Mode().IsDir() {
		fmt.Println("The given path is a directory")
		return
	}

	// Check if the file has read permissions
	if fileInfo.Mode().Perm().IsReadable() {
		fmt.Println("File has read permissions")
	} else {
		fmt.Println("File does not have read permissions")
	}
}
