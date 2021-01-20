package graph

import (
	"fmt"
	"strings"
)

// Graph represents a graph data structure
type Graph struct {
	adjacency map[string][]Edge
}

// Edge represents an edge between 2 nodes
type Edge struct {
	node   string
	weight int
}

// NewGraph creates a New Graph
func NewGraph() Graph {
	return Graph{
		adjacency: make(map[string][]Edge),
	}
}

// AddVertex adds a vertex to a graph
// if vertex alredy added it is ignored
func (g *Graph) AddVertex(vertex string) bool {
	if _, ok := g.adjacency[vertex]; ok {
		fmt.Printf("vertex %v already exists\n", vertex)
		return false
	}

	g.adjacency[vertex] = []Edge{}
	return true
}

// AddEdge adds an egde between two vertices
// if vertex does not exists or edge exists it is ignored
func (g *Graph) AddEdge(vertex, node string, weight int) bool {
	if _, ok := g.adjacency[vertex]; !ok {
		fmt.Printf("vertex %v does not exists\n", vertex)
		return false
	}
	e := Edge{node: node, weight: weight}
	if ok := contains(g.adjacency[vertex], e.node); ok {
		fmt.Printf("node %v already exists\n", node)
		return false
	}

	g.adjacency[vertex] = append(g.adjacency[vertex], e)
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
			if !visited[node.node] {
				q = append(q, node.node)
				visited[node.node] = true
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
		if !visited[node.node] {
			g.dfsRecursive(node.node, visited)
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
		var currentNode string
		currentNode, q = q[0], q[1:]
		path = append(path, currentNode)
		edges := g.adjacency[currentNode]
		if contains(edges, secondNode) {
			path = append(path, secondNode)
			fmt.Println(strings.Join(path, "->"))
			return true
		}

		for _, node := range g.adjacency[currentNode] {
			if !visited[node.node] {
				visited[node.node] = true
				q = append(q, node.node)
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

func contains(slice []Edge, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s.node] = struct{}{}
	}

	_, ok := set[item]
	return ok
}
