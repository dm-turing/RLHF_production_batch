package main

import (
	"fmt"
	"os"
)

// FileNotFoundError is a custom error for when a file is not found.
type FileNotFoundError struct {
	Path string
	err  error
}

func (e *FileNotFoundError) Error() string {
	return fmt.Sprintf("file not found: %s (original error: %v)", e.Path, e.err)
}

// IsFileNotFound returns true if the error is a FileNotFoundError.
func IsFileNotFound(err error) bool {
	_, ok := err.(*FileNotFoundError)
	return ok
}

// PermissionError is a custom error for when there is a permission issue.
type PermissionError struct {
	Path string
	err  error
}

func (e *PermissionError) Error() string {
	return fmt.Sprintf("permission denied: %s (original error: %v)", e.Path, e.err)
}

// IsPermissionError returns true if the error is a PermissionError.
func IsPermissionError(err error) bool {
	_, ok := err.(*PermissionError)
	return ok
}

func openFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return &FileNotFoundError{Path: path, err: err}
		} else if os.IsPermission(err) {
			return &PermissionError{Path: path, err: err}
		}
		return err
	}
	defer file.Close()
	return nil
}

func main() {
	filePath := "non_existent_file.txt"
	err := openFile(filePath)
	if err != nil {
		switch err {
		case nil:
			fmt.Println("File opened successfully.")
		case err.(*FileNotFoundError):
			fmt.Println(err)
		case err.(*PermissionError):
			fmt.Println(err)
		default:
			fmt.Printf("An unexpected error occurred: %v\n", err)
		}
	}
}
