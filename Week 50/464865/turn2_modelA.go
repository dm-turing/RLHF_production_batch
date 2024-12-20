package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/spf13/pflag"
)

var (
	// Flags to be defined
	excludeDirs    string
	fileExtensions string
)

func init() {
	// Define command-line flags
	pflag.StringVarP(&excludeDirs, "exclude-dirs", "e", "", "A comma-separated list of directories to exclude (e.g., '.git,.svn')")
	pflag.StringVarP(&fileExtensions, "file-extensions", "f", "", "A comma-separated list of file extensions to include (e.g., '.go,.txt')")
}

func main() {
	// Parse command-line flags
	pflag.Parse()

	// Get the starting directory path (you can change this to any directory)
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run traverse.go <directory-path>")
		return
	}

	dirPath := os.Args[1]

	// Initialize exclusion and inclusion lists
	excludeDirsSet := make(map[string]bool)
	fileExtensionsSet := make(map[string]bool)

	// Create regular expression patterns for quick exclusion/inclusion checks
	var excludePattern *regexp.Regexp
	if excludeDirs != "" {
		excludeDirsSet = strings.SplitMap(excludeDirs, func(s string) string { return strings.ToLower(s) })
		excludePattern = regexp.MustCompile("\\b(" + strings.Join(excludeDirsSet, "|") + ")\\b")
	}

	var includePattern *regexp.Regexp
	if fileExtensions != "" {
		fileExtensionsSet = strings.SplitMap(fileExtensions, func(s string) string { return strings.ToLower(s) })
		includePattern = regexp.MustCompile("\\b(" + strings.Join(fileExtensionsSet, "|") + ")\\b")
	}

	// Traverse the directory
	traverseDir(dirPath, excludePattern, includePattern)
}

func traverseDir(dirPath string, excludePattern *regexp.Regexp, includePattern *regexp.Regexp) {
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
			// Check if the directory should be excluded
			if excludePattern != nil && excludePattern.MatchString(file.Name()) {
				fmt.Printf("- %s (excluded directory)\n", fullPath)
				continue
			}

			fmt.Printf("- %s (directory)\n", fullPath)
			// Recursively call traverseDir on subdirectories
			traverseDir(fullPath, excludePattern, includePattern)
		} else {
			// Check if the file extension should be included
			if includePattern != nil && !includePattern.MatchString(filepath.Ext(file.Name())) {
				fmt.Printf("- %s (excluded file)\n", fullPath)
				continue
			}

			fmt.Printf("- %s (file)\n", fullPath)
		}
	}
}
