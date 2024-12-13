package main

import (
	"flag"
	"fmt"
	"sync"
	"testing"
)

var benchmarkType = flag.String("type", "mutex", "Type of benchmark: mutex or channel")

type Counter struct {
	value int
	mutex sync.Mutex
}

func (c *Counter) Increment() {
	c.mutex.Lock()
	c.value++
	c.mutex.Unlock()
}

func (c *Counter) Get() int {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.value
}

func TestChannelIncrement(b *testing.B) {
	if *benchmarkType != "channel" {
		b.Skip("Not testing channel")
		return
	}

	ch := make(chan bool, 1)
	var count int

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			ch <- true
			count++
			<-ch
		}
	})

	fmt.Println("Channel Increment Final Count:", count)
}

func TestMutexIncrement(b *testing.B) {
	if *benchmarkType != "mutex" {
		b.Skip("Not testing mutex")
		return
	}

	var counter Counter

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			counter.Increment()
		}
	})

	fmt.Println("Mutex Increment Final Count:", counter.Get())
}

func main() {
	flag.Parse()
	testing.Main()
}
