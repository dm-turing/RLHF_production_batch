package main

import (
	"fmt"
	"os"
)

func checkFileExistenceAndPermissions(path string) (bool, error) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false, fmt.Errorf("file or directory does not exist: %s", path)
	}
	if err != nil {
		return false, err // an error occurred other than not found
	}
	// File or directory exists, we can check permissions
	return true, nil
}

func main() {
	path := "example.txt" // Change this to your file/directory path

	// Check if the file/directory exists and check permissions
	exists, err := checkFileExistenceAndPermissions(path)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if exists {
		fmt.Printf("File or directory '%s' exists.\n", path)

		// Now you can safely perform I/O operations
		file, err := os.Open(path) // We can open the file since it exists
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer file.Close() // Ensure the file is closed after we're done

		// File processing logic goes here
		// Example: Read file content
		contents := make([]byte, 1024)
		_, err = file.Read(contents)
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}
		fmt.Println("File Contents:", string(contents))
	} else {
		fmt.Printf("File or directory '%s' does not exist.\n", path)
	}
}
