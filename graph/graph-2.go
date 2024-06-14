package main

import "fmt"

type graph struct {
	Vertices map[int][]*Vertex
}

type Vertex struct {
	val int
}

func (g *graph) addVertex(key int) {
	if _, ok := (*g).Vertices[key]; !ok {
		(*g).Vertices[key] = make([]*Vertex, 0)
	}
}

func (g *graph) addEdge(key int, e *Vertex) {
	if _, ok := (*g).Vertices[key]; ok {
		(*g).Vertices[key] = append((*g).Vertices[key], e)
	}

}

func (g *graph) printgraph() {
	for v, edges := range (*g).Vertices {
		fmt.Print("Vertex ", v, ": ")
		for _, e := range edges {
			fmt.Print(e.val, " ")
		}
		fmt.Println()
	}
}

func main() {
	g := &graph{
		Vertices: make(map[int][]*Vertex),
	}

	g.addVertex(0)
	g.addVertex(1)
	g.addVertex(2)
	g.addVertex(3)
	g.addVertex(4)

	g.addEdge(0, &Vertex{val: 1})
	g.addEdge(0, &Vertex{val: 2})
	g.addEdge(1, &Vertex{val: 3})
	g.addEdge(2, &Vertex{val: 4})
	g.addEdge(4, &Vertex{val: 3})
	g.addEdge(3, &Vertex{val: 4})

	g.printgraph()
}
