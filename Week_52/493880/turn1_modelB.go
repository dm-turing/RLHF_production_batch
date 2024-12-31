package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type StockPriceObserver interface {
	OnPriceChange(float64)
}

type StockMarket struct {
	observers []StockPriceObserver
	price     float64
	mu        sync.Mutex
}

func (sm *StockMarket) AddObserver(observer StockPriceObserver) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.observers = append(sm.observers, observer)
}

func (sm *StockMarket) RemoveObserver(observer StockPriceObserver) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	for i, o := range sm.observers {
		if o == observer {
			sm.observers = append(sm.observers[:i], sm.observers[i+1:]...)
			return
		}
	}
}

func (sm *StockMarket) StartSimulation() {
	go func() {
		for {
			time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
			sm.mu.Lock()
			newPrice := sm.price + (rand.Float64()-0.5)*2.0
			sm.price = newPrice
			sm.mu.Unlock()

			sm.NotifyObservers()
		}
	}()
}

func (sm *StockMarket) NotifyObservers() {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	for _, observer := range sm.observers {
		observer.OnPriceChange(sm.price)
	}
}

type Investor struct {
	name string
}

func (i *Investor) OnPriceChange(price float64) {
	fmt.Printf("Investor %s: Stock price changed to %.2f\n", i.name, price)
}

func main() {
	stockMarket := &StockMarket{
		price: 100.0,
	}

	investor1 := &Investor{name: "Alice"}
	investor2 := &Investor{name: "Bob"}

	stockMarket.AddObserver(investor1)
	stockMarket.AddObserver(investor2)

	stockMarket.StartSimulation()

	time.Sleep(10 * time.Second) // Simulate the market for 10 seconds

	stockMarket.RemoveObserver(investor2)

	time.Sleep(5 * time.Second) // Simulate the market for another 5 seconds
}
