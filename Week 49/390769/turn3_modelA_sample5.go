package main

import (
	"fmt"
	"os"

	"go.uber.org/zap"
)

// Function that simulates a failure
func ReadFile(filename string) error {
	_, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err) // Wrapping the original error
	}
	return nil
}

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	err := ReadFile("file.txt")
	if err != nil {
		logger.Error("Error reading file", zap.String("filename", "file.txt"), zap.Error(err))
	}
}
