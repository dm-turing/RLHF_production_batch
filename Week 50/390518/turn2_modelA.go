package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/sirupsen/logrus"
)

// Message represents the message to be published.
type Message struct {
	Content string
}

// PubSub is the main structure managing subscribers.
type PubSub struct {
	subscribers []chan Message
	mu          sync.Mutex
}

// NewPubSub initializes a new PubSub instance.
func NewPubSub() *PubSub {
	return &PubSub{
		subscribers: []chan Message{},
	}
}

// Subscribe adds a new subscriber channel.
func (ps *PubSub) Subscribe() chan Message {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	subscriberChan := make(chan Message)
	ps.subscribers = append(ps.subscribers, subscriberChan)

	logrus.WithField("subscriber", fmt.Sprintf("%p", subscriberChan)).Info("New subscriber added.")
	return subscriberChan
}

// Publish sends a message to all subscribers.
func (ps *PubSub) Publish(msg Message) {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	logrus.WithField("message", msg.Content).Info("Publishing message.")

	for _, subscriber := range ps.subscribers {
		go func(sub chan Message) {
			defer func() {
				if r := recover(); r != nil {
					logrus.WithFields(logrus.Fields{
						"error":      fmt.Sprintf("%v", r),
						"subscriber": fmt.Sprintf("%p", sub),
					}).Error("Recovered in Publish - panic occurred.")
				}
			}()

			select {
			case sub <- msg:
				logrus.WithField("subscriber", fmt.Sprintf("%p", sub)).Info("Message sent to subscriber.")
			default:
				logrus.WithField("subscriber", fmt.Sprintf("%p", sub)).Warn("Subscriber is not ready to receive message; skipping.")
			}
		}(subscriber)
	}
}

func subscriber(id int, sub chan Message, wg *sync.WaitGroup) {
	defer wg.Done()

	for msg := range sub {
		logrus.WithFields(logrus.Fields{
			"subscriber": id,
			"message":    msg.Content,
		}).Info("Received message.")
	}
}

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true}) // Set a simple and readable log format

	pubSub := NewPubSub()
	var wg sync.WaitGroup

	const numSubscribers = 3

	// Create subscribers
	for i := 0; i < numSubscribers; i++ {
		wg.Add(1)
		subscriberChan := pubSub.Subscribe()
		go subscriber(i+1, subscriberChan, &wg)
	}

	// Publish messages
	messages := []Message{
		{Content: "Hello, World!"},
		{Content: "Pub-Sub pattern in Go."},
		{Content: "Go is awesome!"},
	}

	// Listen for termination signals gracefully
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		logrus.Info("Received termination signal. Closing subscribers...")
		for _, subscriber := range pubSub.subscribers {
			close(subscriber) // Close each subscriber channel
		}
	}()

	for _, msg := range messages {
		pubSub.Publish(msg)
	}

	wg.Wait() // Wait for all subscribers to finish processing
	logrus.Info("All subscribers have finished processing.")
}
