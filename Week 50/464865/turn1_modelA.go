package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func traverseDir(dirPath string) {
	// Open the directory
	dir, err := os.Open(dirPath)
	if err != nil {
		fmt.Printf("Error opening directory %s: %v\n", dirPath, err)
		return
	}
	defer dir.Close()

	// Iterate over the directory entries
	files, err := dir.Readdir(-1)
	if err != nil {
		fmt.Printf("Error reading directory %s: %v\n", dirPath, err)
		return
	}

	// Print the current directory name
	fmt.Printf("=== %s\n", dirPath)

	// Iterate over each file or subdirectory
	for _, file := range files {
		// Create the full path for the current entry
		fullPath := filepath.Join(dirPath, file.Name())

		// Determine if the entry is a file or directory
		if file.IsDir() {
			fmt.Printf("- %s (directory)\n", fullPath)
			// Recursively call traverseDir on subdirectories
			traverseDir(fullPath)
		} else {
			fmt.Printf("- %s (file)\n", fullPath)
		}
	}
}

func main() {
	// Get the starting directory path (you can change this to any directory)
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run traverse.go <directory-path>")
		return
	}

	dirPath := os.Args[1]

	// Traverse the directory
	traverseDir(dirPath)
}
