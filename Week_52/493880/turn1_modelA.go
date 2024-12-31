package main

import (
	"fmt"
	"sync"
	"time"
)

// Subject interface representing a log publisher
type Subject interface {
	Subscribe(observer <-chan string)
	Unsubscribe(observer <-chan string)
	Notify(msg string)
}

// LogPublisher struct implements the Subject interface
type LogPublisher struct {
	observers []<-chan string
	wrmtx     sync.Mutex
}

func (pub *LogPublisher) Subscribe(observer <-chan string) {
	pub.wrmtx.Lock()
	pub.observers = append(pub.observers, observer)
	pub.wrmtx.Unlock()
}

func (pub *LogPublisher) Unsubscribe(observer <-chan string) {
	pub.wrmtx.Lock()
	defer pub.wrmtx.Unlock()
	for i, c := range pub.observers {
		if c == observer {
			pub.observers = append(pub.observers[:i], pub.observers[i+1:]...)
			return
		}
	}
}

func (pub *LogPublisher) Notify(msg string) {
	for _, c := range pub.observers {
		select {
		case c <- msg:

		default:
			// Handle the case where the observer's channel is full or closed
		}
	}
}

// Observer interface representing logger services
type Observer interface {
	Log(msg string)
}

// ConsoleLogger struct implements the Observer interface
type ConsoleLogger struct {
	ch <-chan string
}

func NewConsoleLogger(ch <-chan string) *ConsoleLogger {
	return &ConsoleLogger{ch}
}

func (cl *ConsoleLogger) Log(msg string) {
	go func() {
		select {
		case <-cl.ch:
			fmt.Printf("ConsoleLogger: %s\n", msg)
		default:
			// Handle logs that can't be processed immediately
		}
	}()
}

// FileLogger struct implements the Observer interface
type FileLogger struct {
	ch <-chan string
}

func NewFileLogger(ch <-chan string) *FileLogger {
	return &FileLogger{ch}
}

func (fl *FileLogger) Log(msg string) {
	go func() {
		select {
		case <-fl.ch:
			fmt.Printf("FileLogger: %s\n", msg)
		default:
			// Handle logs that can't be processed immediately
		}
	}()
}

func main() {
	// Create the subject (log publisher)
	publisher := &LogPublisher{}

	// Create observer channels
	consoleCh := make(chan string, 1)
	fileCh := make(chan string, 1)

	// Create observer instances
	consoleLogger := NewConsoleLogger(consoleCh)
	fileLogger := NewFileLogger(fileCh)

	// Subscribe observers to the publisher
	publisher.Subscribe(consoleCh)
	publisher.Subscribe(fileCh)

	// Publish a log message
	logMsg := "This is a log message"
	publisher.Notify(logMsg)

	// Clean up and unsubscribe observers
	time.Sleep(time.Second) // Ensure logs are processed
	publisher.Unsubscribe(consoleCh)
	publisher.Unsubscribe(fileCh)

	close(consoleCh)
	close(fileCh)
}
