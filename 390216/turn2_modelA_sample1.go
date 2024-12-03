package main

import (
	"fmt"
	"sync"
)

type Node struct {
	id        int
	neighbors []*Node
}

func NewNode(id int) *Node {
	return &Node{
		id:        id,
		neighbors: make([]*Node, 0),
	}
}

func (n *Node) AddNeighbor(neighbor *Node) {
	n.neighbors = append(n.neighbors, neighbor)
}

func concurrentBFS(startNode *Node, fn func(*Node)) {
	var visited map[int]bool
	var wg sync.WaitGroup

	visited = make(map[int]bool)

	var traverse func(*Node)

	traverse = func(node *Node) {
		defer wg.Done()

		if _, found := visited[node.id]; found {
			return
		}

		visited[node.id] = true
		fn(node)

		for _, neighbor := range node.neighbors {
			wg.Add(1)
			go traverse(neighbor)
		}
	}

	wg.Add(1)
	traverse(startNode)
	wg.Wait()
}

func main() {
	// Create the graph
	n1 := NewNode(1)
	n2 := NewNode(2)
	n3 := NewNode(3)
	n4 := NewNode(4)

	n1.AddNeighbor(n2)
	n1.AddNeighbor(n3)
	n2.AddNeighbor(n4)

	// Function to process each node
	processNode := func(node *Node) {
		fmt.Println("Processing node:", node.id)
	}

	// Start concurrent BFS
	concurrentBFS(n1, processNode)
}
