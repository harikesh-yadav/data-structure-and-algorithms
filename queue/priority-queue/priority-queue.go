package main

import (
	"container/heap"
	"fmt"
)

type item struct {
	vertex   int
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
	i := n.(*item)
	size := len(*pq)
	i.index = size
	*pq = append(*pq, i)
}

func (pq *priorityQueue) Pop() any {
	size := len(*pq)
	old := *pq
	item := old[size-1]
	item.index = -1
	old[size-1] = nil
	old = old[:size-1]
	(*pq) = old[0 : size-1]
	return item
}

func main() {

	pq := priorityQueue{}

	pq.Push(&item{vertex: 1, distance: 5, index: 0})
	pq.Push(&item{vertex: 2, distance: 4, index: 1})
	pq.Push(&item{vertex: 4, distance: 2, index: 3})
	pq.Push(&item{vertex: 5, distance: 1, index: 4})
	pq.Push(&item{vertex: 3, distance: 3, index: 2})

	heap.Init(&pq)

	for pq.Len() > 0 {
		i := heap.Pop(&pq).(*item)
		fmt.Println("Vertex:", i.vertex, " Distance:", i.distance)
	}

}
