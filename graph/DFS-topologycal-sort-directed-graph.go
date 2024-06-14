package main

import "fmt"

var visited map[int]bool
var stack []int

func topologySortDirectedGraph(g *graph) {
	if len(g.vertices) == 0 {
		return
	}
	visited = make(map[int]bool)
	stack = make([]int, 0)

	for v := range g.vertices {
		fmt.Println(v, ":")
		if !visited[v] {
			topologySortDirectedGraphHelper((*g).vertices, v, -1)
		}
	}
	fmt.Println(stack)
}

func topologySortDirectedGraphHelper(vertices map[int][]*vertex, start, parent int) {
	visited[start] = true
	for _, neighbor := range vertices[start] {
		if !visited[neighbor.val] {
			topologySortDirectedGraphHelper(vertices, neighbor.val, start)
		}
	}
	stack = append(stack, start)
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
	g.addvertex(5)

	g.addEdge(5, &vertex{val: 0})
	g.addEdge(5, &vertex{val: 2})

	g.addEdge(4, &vertex{val: 0})
	g.addEdge(4, &vertex{val: 1})

	g.addEdge(2, &vertex{val: 3})
	g.addEdge(3, &vertex{val: 1})

	g.printgraph()

	fmt.Println()

	topologySortDirectedGraph(g)

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
		fmt.Print("vertex [", v, "]-> ")
		for _, e := range edges {
			fmt.Print(e.val, " ")
		}
		fmt.Println()
	}
}
