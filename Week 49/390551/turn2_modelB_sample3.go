package main

import (
	"errors"
	"fmt"
	"os"

	"go.uber.org/zap"
)

func openFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return errors.Join(err, fmt.Errorf("failed to open file %s", filename))
	}
	defer func() {
		if err := file.Close(); err != nil {
			zap.S().Error("Error closing file", zap.Error(err))
		}
	}()

	return nil
}

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to create logger:", err)
		os.Exit(1)
	}
	defer logger.Sync() // flushes buffered logs on exit

	zap.S().Info("Application started")

	if err := openFile("non_existent_file.txt"); err != nil {
		zap.S().Error("Error opening file", zap.Error(err))
	}

	zap.S().Info("Application terminated")
}
