package main

import (
	"path/filepath"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
)

func main() {
	// Set up logging with rotation
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{TimestampFormat: time.RFC3339})

	logger.SetHooks(make(logrus.LevelHooks))
	logger.AddHook(&lumberjack.Logger{
		Filename:   "audit.log",
		MaxSize:    500, // MB
		MaxBackups: 3,
		MaxAge:     7, // days
		Compress:   true,
	})

	// Initialize watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		logger.Fatal(err)
	}
	defer watcher.Close()

	// Specify the paths to watch
	paths := []string{"/path/to/watch/1", "/path/to/watch/2"}
	for _, path := range paths {
		err := watcher.Add(path)
		if err != nil {
			logger.Warnf("Unable to add path '%s': %v", path, err)
		} else {
			logger.Infof("Watching path '%s'", path)
		}
	}

	// Start handling events
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				logger.Warn("Watcher closed")
				return
			}
			logger.WithFields(logrus.Fields{
				"EventType": event.Type,
				"Path":      filepath.Clean(event.Name),
			}).Info("File system event detected")

		case err, ok := <-watcher.Errors:
			if !ok {
				logger.Warn("Watcher closed")
				return
			}
			logger.Errorf("Watcher error: %v", err)
		}
	}
}
