package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/rs/zerolog"
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

	log.Info().Str("subscriber_id", fmt.Sprintf("%p", subscriberChan)).Msg("New subscriber added")
	return subscriberChan
}

// Publish sends a message to all subscribers.
func (ps *PubSub) Publish(msg Message) {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	log.Info().Str("message_content", msg.Content).Msg("Publishing message")

	for _, subscriber := range ps.subscribers {
		go func(sub chan Message) {
			defer func() {
				if r := recover(); r != nil {
					log.Error().Err(r).Msg("Recovered in Publish")
				}
			}()
			// Sending message to the subscriber
			select {
			case sub <- msg:
				log.Info().Str("subscriber_id", fmt.Sprintf("%p", sub)).Msg("Message sent to subscriber")
			default:
				log.Warn().Str("subscriber_id", fmt.Sprintf("%p", sub)).Msg("Subscriber is not ready to receive message")
			}
		}(subscriber)
	}
}

// Example subscriber function simulating message processing.
func subscriber(id int, sub chan Message, wg *sync.WaitGroup) {
	defer wg.Done()

	for msg := range sub {
		log.Info().Str("subscriber_id", fmt.Sprintf("%d", id)).Str("message_content", msg.Content).Msg("Subscriber received message")
		// Simulate processing of the message
	}
}

func main() {
	// Configure zerolog
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log.SetOutput(zerolog.ConsoleWriter{})
	pubSub := NewPubSub()

	// Number of subscribers
	const numSubscribers = 3
	var wg sync.WaitGroup

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

	for _, msg := range messages {
		pubSub.Publish(msg)
	}

	// Close subscriber channels after use
	for _, subscriber := range pubSub.subscribers {
		close(subscriber)
	}

	wg.Wait() // Wait for all subscribers to finish processing
	log.Info().Msg("All subscribers have finished processing.")
}
