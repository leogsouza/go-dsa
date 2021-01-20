package graph

import (
	"fmt"
	"strings"
)

// Graph represents a graph data structure
type Graph struct {
	adjacency map[string][]string
}

// NewGraph creates a New Graph
func NewGraph() Graph {
	return Graph{
		adjacency: make(map[string][]string),
	}
}

// AddVertex adds a vertex to a graph
// if vertex alredy added it is ignored
func (g *Graph) AddVertex(vertex string) bool {
	if _, ok := g.adjacency[vertex]; ok {
		fmt.Printf("vertex %v already exists\n", vertex)
		return false
	}

	g.adjacency[vertex] = []string{}
	return true
}

// AddEdge adds an egde between two vertices
// if vertex does not exists or edge exists it is ignored
func (g *Graph) AddEdge(vertex, node string) bool {
	if _, ok := g.adjacency[vertex]; !ok {
		fmt.Printf("vertex %v does not exists\n", vertex)
		return false
	}
	if ok := contains(g.adjacency[vertex], node); ok {
		fmt.Printf("node %v already exists\n", node)
		return false
	}

	g.adjacency[vertex] = append(g.adjacency[vertex], node)
	return true
}

// BFS performs a breadth first search into a graph starting from an initial node
func (g Graph) BFS(start string) {
	visited := g.createVisited()
	var q []string

	visited[start] = true
	q = append(q, start)

	for len(q) > 0 {
		var current string
		current, q = q[0], q[1:]
		fmt.Println("BFS", current)
		for _, node := range g.adjacency[current] {
			if !visited[node] {
				q = append(q, node)
				visited[node] = true
			}
		}
	}
}

// DFS performs a depth first search into a graph starting from an initial node
func (g Graph) DFS(start string) {
	visited := g.createVisited()
	g.dfsRecursive(start, visited)
}

func (g Graph) dfsRecursive(start string, visited map[string]bool) {
	visited[start] = true
	fmt.Println("DFS", start)

	for _, node := range g.adjacency[start] {
		if !visited[node] {
			g.dfsRecursive(node, visited)
		}
	}
}

// CreatePath prints the path between all nodes in a graph
func (g Graph) CreatePath(firstNode, secondNode string) bool {
	visited := g.createVisited()

	var (
		path []string
		q    []string
	)

	q = append(q, firstNode)
	visited[firstNode] = false

	for len(q) > 0 {
		currentNode, q := q[0], q[1:]
		path = append(path, currentNode)
		edges := g.adjacency[currentNode]
		if contains(edges, secondNode) {
			path = append(path, secondNode)
			fmt.Println(strings.Join(path, "->"))
			return true
		}

		for _, node := range g.adjacency[currentNode] {
			if !visited[node] {
				visited[node] = true
				q = append(q, node)
			}
		}
	}

	fmt.Println("no link found")
	return false
}

func (g Graph) createVisited() map[string]bool {
	visited := make(map[string]bool, len(g.adjacency))
	for key := range g.adjacency {
		visited[key] = false
	}

	return visited
}

func contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}
