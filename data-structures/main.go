package main

import (
	"fmt"

	"github.com/leogsouza/go-dsa/data-structures/graph"
)

func main() {
	fmt.Println("Graph DFS and BFS Implementation")
	fmt.Println("DFS: Depth First Search")
	fmt.Println("BFS: Breadth First Search")

	g := graph.NewGraph()

	g.AddVertex("1")
	g.AddVertex("2")
	g.AddVertex("3")
	g.AddVertex("4")
	g.AddVertex("5")
	g.AddVertex("6")

	g.AddEdge("1", "2", 6)
	g.AddEdge("2", "3", 6)
	g.AddEdge("3", "4", 6)
	g.AddEdge("1", "5", 6)

	//g.DFS("B")
	g.CreatePath("1", "6")
	g.BFS("1")

}
