package main

import "fmt"

func topoLogicalSortACyclicGraph(g *graph) {

	inDegree := make([]int, len(g.vertices))
	for _, edges := range g.vertices {
		for _, e := range edges {
			inDegree[e]++
		}
	}
	queue := make([]int, 0)

	for v, c := range inDegree {
		if c == 0 {
			queue = append(queue, v)
		}
	}

	result := make([]int, 0)
	for len(queue) > 0 {
		front := queue[0]
		result = append(result, front)
		queue = queue[1:]
		for _, e := range g.vertices[front] {
			inDegree[e]--
			if inDegree[e] == 0 {
				queue = append(queue, e)
			}
		}
	}

	if len(result) != len(g.vertices) {
		fmt.Println("Graph is cyclic")
	}

	for _, v := range result {
		fmt.Print(v, " ")
	}

}

func main() {
	g := &graph{vertices: make(map[int][]int)}
	g.vertices[0] = []int{1}
	g.vertices[1] = []int{2}
	g.vertices[2] = []int{}
	g.vertices[3] = []int{2, 4}
	g.vertices[4] = []int{}

	topoLogicalSortACyclicGraph(g)

}

type graph struct {
	vertices map[int][]int
}
