package main

import (
	"log"
	"os"

	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("failed to create zap logger: %v", err)
	}
	defer logger.Sync() // flushes buffer

	// Use zap sugar for a simpler log syntax
	sugar := logger.Sugar()

	sugar.Info("Starting application")

	os.Exit(0)
}
