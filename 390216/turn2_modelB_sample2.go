package main

import (
	"container/heap"
	"fmt"
	"math"
)

// PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

// Item is something we manage in a priority queue.
type Item struct {
	node     int
	priority int // The priority of the item in the queue.
	index    int // The index of the item in the heap.
}

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// Graph structure for Dijkstra's algorithm
type Graph struct {
	nodes map[int]map[int]int
}

// NewGraph initializes a new Graph
func NewGraph() *Graph {
	return &Graph{
		nodes: make(map[int]map[int]int),
	}
}

// AddEdge adds an edge with weight to the graph
func (g *Graph) AddEdge(from, to, weight int) {
	if _, ok := g.nodes[from]; !ok {
		g.nodes[from] = make(map[int]int)
	}
	g.nodes[from][to] = weight
}

// Dijkstra finds the shortest path from a start node to all other nodes
func (g *Graph) Dijkstra(start int) map[int]int {
	dist := make(map[int]int)
	for node := range g.nodes {
		dist[node] = math.MaxInt64
	}
	dist[start] = 0

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Item{node: start, priority: 0})

	for pq.Len() > 0 {
		item := heap.Pop(pq).(*Item)
		currentNode := item.node

		for neighbor, weight := range g.nodes[currentNode] {
			if alt := dist[currentNode] + weight; alt < dist[neighbor] {
				dist[neighbor] = alt
				heap.Push(pq, &Item{node: neighbor, priority: alt})
			}
		}
	}

	return dist
}

func main() {
	graph := NewGraph()
	graph.AddEdge(1, 2, 1)
	graph.AddEdge(1, 3, 4)
	graph.AddEdge(2, 3, 2)
	graph.AddEdge(3, 4, 1)

	distances := graph.Dijkstra(1)
	fmt.Println("Shortest distances from node 1:", distances)
}
