package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

// Message represents a message that will be published and subscribed to
type Message struct {
	ID   int
	Data string
}

// Publisher publishes messages to a channel
type Publisher struct {
	msgChan chan Message
	wg      *sync.WaitGroup
	ctx     context.Context
	cancel  context.CancelFunc
}

// NewPublisher creates a new publisher
func NewPublisher() *Publisher {
	return &Publisher{
		msgChan: make(chan Message),
		wg:      &sync.WaitGroup{},
		ctx:     context.Background(),
		cancel:  func() {},
	}
}

// Start starts the publisher
func (p *Publisher) Start() {
	p.ctx, p.cancel = context.WithCancel(p.ctx)
	p.wg.Add(1)

	go func() {
		defer p.wg.Done()
		for {
			select {
			case <-p.ctx.Done():
				log.Println("Publisher shutting down")
				return
			case msg := <-p.msgChan:
				log.Printf("Publisher: Publishing message %+v\n", msg)
				time.Sleep(time.Second / 2) // Simulate some work
			}
		}
	}()
}

// Publish sends a message to the publisher
func (p *Publisher) Publish(msg Message) error {
	select {
	case <-p.ctx.Done():
		return fmt.Errorf("publisher is shutting down")
	default:
		p.msgChan <- msg
		log.Printf("Published message %+v\n", msg)
	}
	return nil
}

// Stop stops the publisher
func (p *Publisher) Stop() {
	p.cancel()
	p.wg.Wait()
}

// Subscriber subscribes to messages from a publisher
type Subscriber struct {
	publisher *Publisher
	wg        *sync.WaitGroup
	ctx       context.Context
	cancel    context.CancelFunc
}

// NewSubscriber creates a new subscriber
func NewSubscriber(pub *Publisher) *Subscriber {
	return &Subscriber{
		publisher: pub,
		wg:        &sync.WaitGroup{},
		ctx:       context.Background(),
		cancel:    func() {},
	}
}

// Start starts the subscriber
func (s *Subscriber) Start() {
	s.ctx, s.cancel = context.WithCancel(s.ctx)
	s.wg.Add(1)

	go func() {
		defer s.wg.Done()
		for {
			select {
			case <-s.ctx.Done():
				log.Println("Subscriber shutting down")
				return
			case msg := <-s.publisher.msgChan:
				log.Printf("Subscriber: Received message %+v\n", msg)
				time.Sleep(time.Second / 2) // Simulate some work
			}
		}
	}()
}

// Stop stops the subscriber
func (s *Subscriber) Stop() {
	s.cancel()
	s.wg.Wait()
}

func main() {
	publisher := NewPublisher()
	subscriber := NewSubscriber(publisher)

	go publisher.Start()
	go subscriber.Start()

	time.Sleep(time.Second)

	err := publisher.Publish(Message{ID: 1, Data: "Hello, World!"})
	if err != nil {
		log.Fatalf("Error publishing message: %v\n", err)
	}

	err = publisher.Publish(Message{ID: 2, Data: "Goodbye!"})
	if err != nil {
		log.Fatalf("Error publishing message: %v\n", err)
	}

	time.Sleep(2 * time.Second)

	publisher.Stop()
	subscriber.Stop()
}
