package main

import (
	"container/heap"
	"fmt"
	"math"
)

/***
Problem:https://www.geeksforgeeks.org/dijkstras-shortest-path-algorithm-greedy-algo-7/

Time Complexity: O(E * logV), Where E is the number of edges and V is the number of vertices.
Auxiliary Space: O(V)

***/

type vertex struct {
	value    int
	distance int
}

type graph struct {
	//  [0]=[]{vertex}
	vertices map[int][]vertex
}

func dijkstraAlgorithmOneToAll(g *graph) {
	if len(g.vertices) == 0 {
		return
	}

	visited := make(map[int]bool)
	paths := make(map[int]int)

	for i := 0; i < 6; i++ {
		paths[i] = math.MaxInt
	}

	pq := priorityQueue{}
	pq.Push(&item{value: 0, distance: 0, index: 0})

	for pq.Len() > 0 {
		front := heap.Pop(&pq).(*item)
		if !visited[front.value] {
			visited[front.value] = true
			if paths[front.value] > front.distance {
				paths[front.value] = front.distance
			}

			for _, v := range g.vertices[front.value] {
				if !visited[v.value] {
					pq.Push(&item{value: v.value, distance: v.distance + paths[front.value], index: pq.Len()})
				}
			}
			heap.Init(&pq)
		}
	}

	for k, v := range paths {
		fmt.Println(k, "-", v)
	}
}

func main() {
	g := &graph{
		vertices: make(map[int][]vertex),
	}

	g.vertices[0] = []vertex{
		vertex{value: 2, distance: 4},
		vertex{value: 3, distance: 2},
	}

	g.vertices[1] = []vertex{
		vertex{value: 2, distance: 9},
		vertex{value: 4, distance: 1},
	}

	g.vertices[2] = []vertex{
		vertex{value: 0, distance: 4},
		vertex{value: 1, distance: 9},
		vertex{value: 3, distance: 3},
		vertex{value: 5, distance: 1},
	}

	g.vertices[3] = []vertex{
		vertex{value: 0, distance: 2},
		vertex{value: 2, distance: 3},
		vertex{value: 5, distance: 1},
	}

	g.vertices[4] = []vertex{
		vertex{value: 1, distance: 1},
		vertex{value: 5, distance: 5},
	}

	g.vertices[5] = []vertex{
		vertex{value: 3, distance: 1},
		vertex{value: 2, distance: 1},
		vertex{value: 4, distance: 4},
	}

	dijkstraAlgorithmOneToAll(g)

}

type item struct {
	value    int
	distance int
	index    int
}

type priorityQueue []*item

func (pq priorityQueue) Len() int {
	return len(pq)
}
func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].distance < pq[j].distance
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *priorityQueue) Push(n any) {
	size := len(*pq)
	item := n.(*item)
	item.index = size
	*pq = append(*pq, item)
}

func (pq *priorityQueue) Pop() any {
	size := len(*pq)
	old := *pq
	item := old[size-1]
	item.index = -1
	old[size-1] = nil
	old = old[:size-1]
	*pq = old[0 : size-1]
	return item
}
