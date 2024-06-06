package main

/***
Problem: https://www.geeksforgeeks.org/print-nodes-distance-k-given-node-binary-tree/

time complexity:  O(n)
space complexity O(1)
***/

import (
	"fmt"
)

func findTargetNode(n *node, target int) *node {
	if n == nil {
		return nil
	}
	if n.data == target {
		return n
	}
	l := findTargetNode(n.left, target)
	if l != nil {
		return l
	}
	return findTargetNode(n.right, target)
}

func findParent(n, prev *node) {
	if n == nil {
		return
	}
	parent[n.data] = prev
	prev = n
	findParent(n.left, prev)
	findParent(n.right, prev)
}

func PrintDownFromTargetNode(n *node, k int) {
	if n == nil {
		return
	}
	if visited[n.data] {
		return
	}

	if k == 0 {
		fmt.Print(n.data, " ")
	}
	PrintDownFromTargetNode(n.left, k-1)
	PrintDownFromTargetNode(n.right, k-1)
}

func printTopFromTargetNode(n *node, k int) {
	if n == nil {
		return
	}
	if visited[n.data] {
		return
	}
	visited[n.data] = true

	if k == 0 {
		fmt.Print(n.data, " ")
		return
	}
	printTopFromTargetNode(parent[n.data], k-1)
	PrintDownFromTargetNode(n.left, k-1)
	PrintDownFromTargetNode(n.right, k-1)

}

func printKthDistanceNodesFromGivenNode(n *node, targetNode, k int) {
	if n == nil {
		return
	}
	target := findTargetNode(n, targetNode)
	if target == nil {
		return
	}
	findParent(n, nil)

	visited[target.data] = true

	printTopFromTargetNode(parent[target.data], k-1)

	PrintDownFromTargetNode(target.left, k-1)
	PrintDownFromTargetNode(target.right, k-1)
}

var parent map[int]*node
var visited map[int]bool

func main() {

	n := &node{data: 20}
	n.left = &node{data: 8}
	n.right = &node{data: 22}

	n.left.left = &node{data: 4}
	n.left.right = &node{data: 12}

	n.left.right.left = &node{data: 10}
	n.left.right.right = &node{data: 14}

	targetNode := 8

	k := 1

	parent = make(map[int]*node)
	visited = make(map[int]bool)
	printKthDistanceNodesFromGivenNode(n, targetNode, k)

}

type node struct {
	data  int
	left  *node
	right *node
}
