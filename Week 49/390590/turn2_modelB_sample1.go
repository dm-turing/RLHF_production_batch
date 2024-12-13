package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numProducers = 3
	numConsumers = 2
	bufferSize   = 10
)

var wg sync.WaitGroup

func producer(ch chan int) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		value := rand.Intn(100)
		ch <- value
		fmt.Printf("Producer %d sent: %d\n", i, value)
		time.Sleep(time.Millisecond * 75)
	}
}

func consumer(ch chan int) {
	defer wg.Done()
	for value := range ch {
		fmt.Printf("Consumer received: %d\n", value)
		time.Sleep(time.Millisecond * 20)
	}
}

func main() {
	ch := make(chan int, bufferSize)

	wg.Add(numProducers)
	for i := 0; i < numProducers; i++ {
		go producer(ch)
	}

	for i := 0; i < numConsumers; i++ {
		wg.Add(numConsumers)
		go consumer(ch)
	}

	wg.Wait()
	close(ch)
}
