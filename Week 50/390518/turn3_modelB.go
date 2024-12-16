package main

import (
	"fmt"
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

	return subscriberChan
}

// Publish sends a message to all subscribers.
func (ps *PubSub) Publish(msg Message) {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	for _, subscriber := range ps.subscribers {
		go func(sub chan Message) {
			defer func() {
				if r := recover(); r != nil {
					log.Printf("Recovered in subscriber: %v", r)
					// Consider logging a stack trace using runtime.Stack(), closing the channel, etc.
				}
			}()

			sub <- msg
		}(subscriber)
	}
}

// Example subscriber function with logic that may panic.
func subscriber(id int, sub chan Message, wg *sync.WaitGroup) {
	defer wg.Done()

	for msg := range sub {
		// Logic that could potentially panic
		// Unexpected element type handling
		elem, j := msg.Content, make([]int, 0)
		if i, err := fmt.Sscanf(elem, "%d", &j); err != nil || i != 1 {
			panic("Unexpected content format: " + elem)
		}

		log.Printf("Subscriber %d processed message: %s", id, msg.Content)
	}
}

func main() {
	pubSub := NewPubSub()
	var wg sync.WaitGroup

	const numSubscribers = 2

	// Create subscribers
	for i := 0; i < numSubscribers; i++ {
		wg.Add(1)
		go subscriber(i+1, pubSub.Subscribe(), &wg)
	}

	// Publish messages that might trigger a panic in one subscriber
	messages := []Message{
		{Content: "42"},
		{Content: "Invalid Data"}, // This will cause a panic
		{Content: "1337"},
	}

	for _, msg := range messages {
		pubSub.Publish(msg)
	}

	wg.Wait() // Wait for all subscribers to finish
	log.Println("All subscribers have finished processing.")
}
