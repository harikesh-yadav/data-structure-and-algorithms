package main

import "fmt"

func dfsTraversal(g *graph, start int) {
	if len(g.vertices) == 0 {
		return
	}
	visited := map[int]bool{}
	dfsTraversalUtil(g, start, &visited)

}

func dfsTraversalUtil(g *graph, vertex int, visited *map[int]bool) {
	(*visited)[vertex] = true
	fmt.Print(vertex, " ")
	for _, edge := range g.vertices[vertex] {
		if !(*visited)[edge.val] {
			dfsTraversalUtil(g, edge.val, visited)
		}
	}
}

func main() {
	g := &graph{
		vertices: make(map[int][]*vertex),
	}

	g.addvertex(0)
	g.addvertex(1)
	g.addvertex(2)
	g.addvertex(3)
	g.addvertex(4)

	g.addEdge(0, &vertex{val: 2})
	g.addEdge(0, &vertex{val: 1})
	g.addEdge(1, &vertex{val: 3})
	g.addEdge(1, &vertex{val: 0})
	g.addEdge(2, &vertex{val: 4})
	g.addEdge(2, &vertex{val: 0})
	g.addEdge(4, &vertex{val: 3})
	g.addEdge(4, &vertex{val: 2})
	g.addEdge(3, &vertex{val: 4})
	g.addEdge(3, &vertex{val: 1})

	//g.printgraph()

	dfsTraversal(g, 0)
}

type graph struct {
	vertices map[int][]*vertex
}

type vertex struct {
	val int
}

func (g *graph) addvertex(key int) {
	if _, ok := (*g).vertices[key]; !ok {
		(*g).vertices[key] = make([]*vertex, 0)
	}
}

func (g *graph) addEdge(key int, edge *vertex) {
	if edges, ok := (*g).vertices[key]; ok {
		(*g).vertices[key] = append(edges, edge)
	}
}

func (g *graph) printgraph() {
	for v, edges := range (*g).vertices {
		fmt.Print("vertex ", v, "-> ")
		for _, e := range edges {
			fmt.Print(e.val, " ")
		}
		fmt.Println()
	}
}
