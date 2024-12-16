package main

import (
	"os"

	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		logger.Fatal(err.Error())
	}
	defer logger.Sync()

	logger.Info("Application started")

	os.Exit(0)
}
