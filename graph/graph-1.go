package main

import "fmt"

type Graph struct {
	Vertices map[int]*Vertex
}

type Vertex struct {
	val   int
	Edges []*Edge
}

type Edge struct {
	Weight int
	Vertex *Vertex
}

func (g *Graph) AddVertex(v *Vertex) {
	(*g).Vertices[v.val] = v
}

func (g *Graph) AddEdge(vertex int, e *Edge) {
	if edges, ok := (*g).Vertices[vertex]; ok {
		edges.Edges = append(edges.Edges, e)
	}
}

func (g *Graph) printGraph() {
	for v, vertex := range (*g).Vertices {
		fmt.Print("Vertex ", v, ":")
		for _, e := range vertex.Edges {
			fmt.Print("Weight :", e.Weight, " Edge:", e.Vertex.val)
		}
		fmt.Println()
	}
}

func main() {
	g := &Graph{
		Vertices: make(map[int]*Vertex),
	}

	g.AddVertex(&Vertex{val: 0})
	g.AddVertex(&Vertex{val: 1})
	g.AddVertex(&Vertex{val: 2})
	g.AddVertex(&Vertex{val: 3})
	g.AddVertex(&Vertex{val: 4})

	g.AddEdge(0, &Edge{Weight: 1, Vertex: &Vertex{val: 1}})
	g.AddEdge(0, &Edge{Weight: 2, Vertex: &Vertex{val: 2}})
	g.AddEdge(1, &Edge{Weight: 3, Vertex: &Vertex{val: 3}})
	g.AddEdge(2, &Edge{Weight: 4, Vertex: &Vertex{val: 4}})
	g.AddEdge(3, &Edge{Weight: 5, Vertex: &Vertex{val: 4}})

	g.printGraph()
}
