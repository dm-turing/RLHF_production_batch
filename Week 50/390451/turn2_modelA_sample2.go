package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	log.WithFields(log.Fields{
		"event": "event_name",
		"topic": "topic_name",
	}).Info("Event occurred")
}
