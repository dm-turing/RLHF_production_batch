package main

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

// Function that simulates a failure
func ReadFile(filename string) error {
	_, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err) // Wrapping the original error
	}
	return nil
}

func main() {
	err := ReadFile("file.txt")
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"filename": "file.txt",
			"error":    err.Error(),
		}).Error("Error reading file")
	}
}
