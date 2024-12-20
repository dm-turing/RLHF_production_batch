package main

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func traverseDir(dirPath string, showFiles, showDirs bool) error {
	// Use filepath.WalkDir to recursively traverse the directory tree
	return filepath.WalkDir(dirPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Print file/directory info based on flags
		if showFiles && !d.IsDir() {
			fmt.Printf("- %s\n", path)
		} else if showDirs && d.IsDir() {
			fmt.Printf("- %s\n", path)
		}
		return nil
	})
}

func main() {
	// Define command-line flags
	var (
		dirPath   = flag.String("dir", ".", "Directory path to traverse")
		recursive = flag.Bool("R", false, "Traverse the directory recursively")
		showFiles = flag.Bool("f", true, "Show files in the output")
		showDirs  = flag.Bool("d", true, "Show directories in the output")
	)

	// Parse command-line flags
	flag.Parse()

	// Get the absolute path of the starting directory
	absDirPath, err := filepath.Abs(*dirPath)
	if err != nil {
		log.Fatalf("Error getting absolute path: %v", err)
	}

	// Set the "recursive" flag to true if -R flag is provided
	if *recursive {
		*showDirs = true
		*showFiles = true
	}

	// Validate the provided directory path
	info, err := os.Stat(absDirPath)
	if err != nil {
		log.Fatalf("Error accessing directory %s: %v", absDirPath, err)
	}

	if !info.IsDir() {
		log.Fatalf("%s is not a directory", absDirPath)
	}

	// Traverse the directory using the flags
	err = traverseDir(absDirPath, *showFiles, *showDirs)
	if err != nil {
		log.Fatalf("Error traversing directory: %v", err)
	}
}
