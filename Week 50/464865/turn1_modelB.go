package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func main() {
	root := "." // Current working directory
	if len(os.Args) > 1 {
		root = os.Args[1] // Use the provided argument as the root directory
	}
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			fmt.Printf("d %s\n", path)
		} else {
			fmt.Printf("- %s\n", path)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("Error walking the path: %v\n", err)
	}
}
