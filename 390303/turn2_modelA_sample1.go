package main

import (
	"context"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var m sync.Mutex

	go func() {
		m.Lock()
		time.Sleep(time.Second * 3) // Simulate long operation
		m.Unlock()
	}()

	go func(ctx context.Context) {
		select {
		case <-ctx.Done():
			print("Deadlock detected or timeout occurred.")
			return
		default:
			m.Lock()
			time.Sleep(time.Second * 3) // Simulate long operation
			m.Unlock()
		}
	}(ctx)
}
