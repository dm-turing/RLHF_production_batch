package main

import (
	"fmt"
	"sync"
)

// GraphNode represents a node in a graph.
type GraphNode struct {
	Value    int
	Children []*GraphNode
}

// concurrentBFS performs a concurrent breadth-first search on a graph.
func concurrentBFS(graph *GraphNode, waitGroup *sync.WaitGroup, visitCh chan<- string) {
	defer waitGroup.Done()

	queue := []*GraphNode{graph}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		// Construct a printable version of the node's value and its children
		visitMessage := fmt.Sprintf("Visiting node %d, Children: %v", node.Value, node.Children)
		visitCh <- visitMessage

		for _, child := range node.Children {
			queue = append(queue, child)
		}
	}
}

func main() {
	// Create a graph
	graph := &GraphNode{
		Value: 1,
		Children: []*GraphNode{
			&GraphNode{Value: 2, Children: []*GraphNode{{Value: 4}}},
			&GraphNode{Value: 3, Children: []*GraphNode{{Value: 5}, {Value: 6}}},
		},
	}

	// Channel to store the visit messages
	visitCh := make(chan string)

	// WaitGroup to ensure the main thread waits for all goroutines to finish
	var waitGroup sync.WaitGroup

	// Start goroutines for concurrent traversal
	for i := 0; i < 3; i++ {
		waitGroup.Add(1)
		go concurrentBFS(graph, &waitGroup, visitCh)
	}

	// Print the visit messages
	for message := range visitCh {
		fmt.Println(message)
	}

	// Wait for all goroutines to finish
	waitGroup.Wait()
}
