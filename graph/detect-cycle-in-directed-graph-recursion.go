package main

import "fmt"

var visited map[int]bool
var recusrsion map[int]bool

func detectCycleInDirectedGraph(g *graph) bool {
	if len(g.vertices) == 0 {
		return false
	}
	visited = make(map[int]bool)
	recusrsion = make(map[int]bool)

	for v := range g.vertices {
		if !visited[v] {
			if detectCycleHelper((*g).vertices, v, -1) {
				return true
			}
		}

	}
	return false
}

func detectCycleHelper(vertices map[int][]*vertex, start, parent int) bool {
	visited[start] = true
	recusrsion[start] = true
	fmt.Print(start, " ")
	for _, neighbor := range vertices[start] {
		if !visited[neighbor.val] {
			if detectCycleHelper(vertices, neighbor.val, start) {
				return true
			}
		} else if recusrsion[neighbor.val] {
			return true
		}
	}
	recusrsion[start] = false
	return false
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

	g.addEdge(0, &vertex{val: 1})

	g.addEdge(2, &vertex{val: 1})
	g.addEdge(2, &vertex{val: 3})

	g.addEdge(3, &vertex{val: 4})

	g.addEdge(4, &vertex{val: 2})

	g.printgraph()
	fmt.Println()

	fmt.Println("IS Graph have Cycle:", detectCycleInDirectedGraph(g))
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
