package main

import (
	"log"
	"sync"
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

	log.Printf("New subscriber added: %p", subscriberChan) // Log new subscriber
	return subscriberChan
}

// Publish sends a message to all subscribers.
func (ps *PubSub) Publish(msg Message) {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	log.Printf("Publishing message: %s", msg.Content) // Log published message

	for _, subscriber := range ps.subscribers {
		go func(sub chan Message) {
			defer func() {
				if r := recover(); r != nil {
					log.Printf("Recovered in Publish: %v", r) // Log panic recovery
				}
			}()
			// Sending message to the subscriber
			select {
			case sub <- msg:
				log.Printf("Message sent to subscriber %p", sub) // Log successful send
			default:
				log.Printf("Subscriber %p is not ready to receive message", sub) // Log not ready
			}
		}(subscriber)
	}
}

// Example subscriber function simulating message processing.
func subscriber(id int, sub chan Message, wg *sync.WaitGroup) {
	defer wg.Done()

	for msg := range sub {
		log.Printf("Subscriber %d received message: %s", id, msg.Content)
		// Simulate processing of the message
	}
}

func main() {
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

	// Close subscriber channels after use (for a real-world scenario)
	for _, subscriber := range pubSub.subscribers {
		close(subscriber)
	}

	wg.Wait() // Wait for all subscribers to finish processing
	log.Println("All subscribers have finished processing.")
}
