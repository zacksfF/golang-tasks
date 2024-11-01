package main

import (
	"container/heap"
	"fmt"
	"math"
)

//Based on this challenge involves finding the Optimal Path
//between two nodes (routers) in a network graph.

// Define the graph structure and preiority queue
type Edge struct {
	to     string
	weight int
}

type Item struct {
	node         string
	latency      int
	IsCompressed bool
}

// Priority Queue Implemenattion for Dijkstra
type PriorityQueue []Item

// A helper functions for maintain the main function
func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].latency < pq[j].latency
}

func (pq *PriorityQueue) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(Item))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}

// // Helper function to calculate minimum of two values
// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }

// find_minimum_latency_path to find the minimmum latency path
func find_minimum_latency_path(graph map[string][]Edge, compressionNodes map[string]bool, source, destination string) int {

	//init the priority queue
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, Item{node: source, latency: 0, IsCompressed: false})

	//init latencies map with infinity values
	latencies := make(map[string]map[bool]int)
	for node := range graph {
		latencies[node] = map[bool]int{false: math.MaxInt64, true: math.MaxInt64}
	}
	latencies[source][false] = 0

	for pq.Len() > 0 {
		current := heap.Pop(pq).(Item)

		// If we reached the destination node, return the latency
		if current.node == destination && (current.latency < latencies[destination][current.IsCompressed]) {
			return current.latency
		}

		// Skip this path if we already found a better one
		if current.latency > latencies[current.node][current.IsCompressed] {
			continue
		}

		for _, edge := range graph[current.node] {
			nextLatency := current.latency + edge.weight

			// Update without compression
			if nextLatency < latencies[edge.to][current.IsCompressed] {
				latencies[edge.to][current.IsCompressed] = nextLatency
				heap.Push(pq, Item{node: edge.to, latency: nextLatency, IsCompressed: current.IsCompressed})
			}

			// Update with compression if available and not used yet
			if !current.IsCompressed && compressionNodes[current.node] {
				compressedLatency := current.latency + edge.weight/2
				if compressedLatency < latencies[edge.to][true] {
					latencies[edge.to][true] = compressedLatency
					heap.Push(pq, Item{node: edge.to, latency: compressedLatency, IsCompressed: true})
				}
			}
		}
	}

	return -1 // Return -1 if there's no path to the destination
}

func main() {
	//usgae
	graph := map[string][]Edge{
		"A": {{"B", 10}, {"C", 20}},
		"B": {{"D", 15}},
		"C": {{"D", 30}},
		"D": {},
	}

	//compression nodes
	compressionNodes := map[string]bool{"B": true, "C": true}

	//Test the function
	source := "A"
	destination := "D"
	result := find_minimum_latency_path(graph, compressionNodes, source, destination)
	fmt.Printf("Minimum total latency from %s to %s: %d\n", source, destination, result)
}
