package main

import (
	"fmt"
	"sync"
)

type Resource struct {
	name string
}

func (r *Resource) Acquire() {
	fmt.Printf("Acquiring resource: %s\n", r.name)
}

func (r *Resource) Release() {
	fmt.Printf("Releasing resource: %s\n", r.name)
}

func useResource(r *Resource) {
	r.Acquire()
	defer r.Release()

	// Use the resource
	fmt.Printf("Using resource: %s\n", r.name)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		res1 := &Resource{name: "Resource1"}
		useResource(res1)
	}()

	go func() {
		defer wg.Done()
		res2 := &Resource{name: "Resource2"}
		useResource(res2)
	}()

	wg.Wait()
}
