package main

import (
	"fmt"
	"sync"
)

// Graph structure
type Graph struct {
	nodes map[int][]int
	mu    sync.RWMutex
}

// NewGraph initializes a new Graph
func NewGraph() *Graph {
	return &Graph{
		nodes: make(map[int][]int),
	}
}

// AddEdge adds an edge to the graph
func (g *Graph) AddEdge(from, to int) {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.nodes[from] = append(g.nodes[from], to)
	g.nodes[to] = append(g.nodes[to], from) // if undirected graph
}

// GetNeighbors safely retrieves neighbors of a given node
func (g *Graph) GetNeighbors(node int) []int {
	g.mu.RLock()
	defer g.mu.RUnlock()
	return g.nodes[node]
}

// ConcurrentGraphTraversal traverses nodes concurrently
func ConcurrentGraphTraversal(g *Graph, startNode int, fn func(int)) {
	var wg sync.WaitGroup
	visited := make(map[int]bool)
	var mu sync.Mutex

	var traverse func(node int)

	traverse = func(node int) {
		defer wg.Done()

		// Lock before accessing the shared visited map
		mu.Lock()
		if visited[node] {
			mu.Unlock()
			return
		}
		visited[node] = true
		mu.Unlock()

		// Perform the function on the node
		fn(node)

		neighbors := g.GetNeighbors(node)
		for _, neighbor := range neighbors {
			wg.Add(1)
			go traverse(neighbor)
		}
	}

	wg.Add(1)
	traverse(startNode)
	wg.Wait()
}

func main() {
	// Create a new graph
	graph := NewGraph()

	// Add edges to the graph
	graph.AddEdge(1, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(2, 4)
	graph.AddEdge(3, 4)
	graph.AddEdge(4, 5)

	// Function to run on each node during traversal
	nodeFn := func(node int) {
		fmt.Println("Visited node:", node)
	}

	// Start concurrent traversal from node 1
	ConcurrentGraphTraversal(graph, 1, nodeFn)
}
