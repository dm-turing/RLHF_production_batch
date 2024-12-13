package main

import (
	"fmt"
	"os"

	"go.uber.org/zap"
)

func createFile(filename string) error {
	// Set up structured logger
	logger, _ := zap.NewProduction()
	defer logger.Sync() // Flush any buffered log entries

	file, err := os.Create(filename)
	if err != nil {
		logger.Error("Failed to create file", zap.String("filename", filename), zap.Error(err))
		return fmt.Errorf("createFile error: %w", err)
	}

	defer func() {
		closeErr := file.Close()
		if closeErr != nil {
			logger.Error("Failed to close file", zap.String("filename", filename), zap.Error(closeErr))
		}
	}()

	if _, err := file.WriteString("Hello, world!"); err != nil {
		logger.Error("Failed to write to file", zap.String("filename", filename), zap.Error(err))
		return fmt.Errorf("writeFile error: %w", err)
	}

	return nil
}

func main() {
	err := createFile("example.txt")
	if err != nil {
		fmt.Println("An error occurred:", err)
		// Log error
	}
}
