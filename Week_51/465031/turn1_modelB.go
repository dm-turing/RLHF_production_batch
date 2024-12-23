package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numProducers = 5
	numConsumers = 3
	bufferSize   = 10
)

var (
	wg      sync.WaitGroup
	dataMap sync.Map
)

func producer(id int) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		randNum := rand.Intn(1000)
		dataMap.Store(randNum, nil)
		fmt.Printf("Producer %d: Produced %d\n", id, randNum)
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	}
}

func consumer(id int) {
	defer wg.Done()
	for {
		// Loop to retrieve data from the map
		dataMap.Range(func(key, value interface{}) bool {
			fmt.Printf("Consumer %d: Consumed %d\n", id, key)
			// Remove the consumed data from the map
			dataMap.Delete(key)
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			return true // Continue the loop
		})
		// If the map is empty, break the loop
		if dataMap.Len() == 0 {
			break
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	wg.Add(numProducers + numConsumers)

	// Start producers
	for i := 0; i < numProducers; i++ {
		go producer(i)
	}

	// Start consumers
	for i := 0; i < numConsumers; i++ {
		go consumer(i)
	}

	wg.Wait()
	fmt.Println("All producers and consumers finished.")
}
