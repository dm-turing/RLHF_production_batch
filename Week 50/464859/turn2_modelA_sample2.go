package main

import (
	"fmt"
	"log"
)

// Logging decorator for a Reader interface method
func LoggingDecorator(r Reader) Reader {
	return &loggingReader{r}
}

type loggingReader struct {
	Reader
}

func (lr *loggingReader) Read(p []byte) (n int, err error) {
	n, err = lr.Reader.Read(p)
	if err != nil {
		log.Printf("Error reading: %v\n", err)
	}
	return n, err
}

// Simple implementation of the Reader interface
type simpleReader struct {
	data []byte
}

func (sr *simpleReader) Read(p []byte) (n int, err error) {
	n, err = copy(p, sr.data)
	sr.data = sr.data[n:]
	return n, err
}

func main() {
	// Create a simple reader
	sr := simpleReader{data: []byte("Hello, world!")}

	// Decorate the reader with the logging decorator
	lr := LoggingDecorator(&sr)

	// Read from the decorated reader
	p := make([]byte, 100)
	n, err := lr.Read(p)
	if err != nil {
		fmt.Printf("Error in main: %v\n", err)
	}
	fmt.Printf("Read: %s\n", string(p[:n]))
}
