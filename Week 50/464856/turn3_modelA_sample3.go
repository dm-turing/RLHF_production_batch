package main

import (
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

func main() {
	logFile := filepath.Join(os.TempDir(), "audit.log")
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(&logrus.FileWriter{Filename: logFile})

	// Log rotation every 1 day
	hook := &logrus.RotateFileHook{
		Filename:   logFile,
		MaxSize:    10, // 10 MB
		MaxBackups: 5,
		MaxAge:     30, // 30 days
	}
	logger.AddHook(hook)

	// ... rest of the code
}
