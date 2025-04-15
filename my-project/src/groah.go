package main

import "fmt"

// Graph structure
type Graph struct {
	Vertices map[int]*Vertex
}

// Vertex structure
type Vertex struct {
	Value int
	Edges map[int]*Edge
}

// Edge structure
type Edge struct {
	Weight int
	Vertex *Vertex
}

// AddVertex adds a vertex to the graph
func (g *Graph) AddVertex(value int) {
	if g.Vertices == nil {
		g.Vertices = make(map[int]*Vertex)
	}
	g.Vertices[value] = &Vertex{Value: value, Edges: make(map[int]*Edge)}
}

// AddEdge adds an edge between two vertices with a specified weight
func (g *Graph) AddEdge(from, to, weight int) {
	fromVertex, fromExists := g.Vertices[from]
	toVertex, toExists := g.Vertices[to]

	if !fromExists || !toExists {
		return // Ensure both vertices exist
	}

	fromVertex.Edges[to] = &Edge{Weight: weight, Vertex: toVertex}

	toVertex.Edges[from] = &Edge{Weight: weight, Vertex: fromVertex}
}

func (g *Graph) PrintGraph() {
	for _, vertex := range g.Vertices {
		fmt.Printf("Vertex %d:\n", vertex.Value)
		for _, edge := range vertex.Edges {
			fmt.Printf("  -> To Vertex %d with weight %d\n", edge.Vertex.Value, edge.Weight)
		}
	}
}
func getNeb(curr int, adjList map[int][]int) []int {
	return adjList[curr]
}
func GraphBFSTraversal(adjlist map[int][]int, start int, end int) {
	queue := []int{start}         // Queue for BFS
	visited := make(map[int]bool) // Track visited nodes
	visited[start] = true

	parent := make(map[int]int) // To reconstruct the shortest path

	for len(queue) > 0 {
		levelSize := len(queue) // Number of nodes in the current level

		for i := 0; i < levelSize; i++ {
			current := queue[0] // Dequeue
			queue = queue[1:]

			// If we reached the destination
			if current == end {
				// Print the shortest path
				path := []int{}
				for node := end; node != start; node = parent[node] {
					path = append([]int{node}, path...) // Prepend to path
				}
				path = append([]int{start}, path...) // Add start node
				fmt.Println("Shortest Path:", path)
				return
			}

			// Visit neighbors
			for _, neighbor := range getNeb(current, adjlist) {
				if !visited[neighbor] {
					visited[neighbor] = true
					parent[neighbor] = current // Track parent for path reconstruction
					queue = append(queue, neighbor)
				}
			}
		}
	}

	fmt.Println("No path found from", start, "to", end)
}

func main() {
	// Create a new graph
	g := &Graph{}

	// Add vertices
	g.AddVertex(1)
	g.AddVertex(2)
	g.AddVertex(3)

	// Add edges
	g.AddEdge(1, 2, 10)
	g.AddEdge(1, 3, 5)
	g.AddEdge(2, 3, 2)
	adjList := map[int][]int{
		1: {1, 3},
		2: {1, 3},
		3: {1, 2},
	}
	GraphBFSTraversal(adjList, 1, 3)
	// Print the graph
	g.PrintGraph()
}
