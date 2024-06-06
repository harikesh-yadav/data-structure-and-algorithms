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

func PrintDownFromTargetNode(n *node, k int) {
	if n == nil {
		return
	}
	if k == 0 {
		fmt.Print(n.data, " ")
	}
	PrintDownFromTargetNode(n.left, k-1)
	PrintDownFromTargetNode(n.right, k-1)

}

func printKthDistanceNodesFromGivenNode(n, target *node, k int) int {
	if n == nil || target == nil {
		return -1
	}

	if n == target {
		PrintDownFromTargetNode(target, k)
		return 0
	}

	ld := printKthDistanceNodesFromGivenNode(n.left, target, k)
	if ld != -1 {
		if ld+1 == k {
			fmt.Print(n.data, " ")
		} else {
			PrintDownFromTargetNode(n.right, k-2-ld)
		}
		return ld + 1
	}

	rd := printKthDistanceNodesFromGivenNode(n.right, target, k)
	if rd != -1 {
		if ld+1 == k {
			fmt.Print(n.data, " ")
		} else {
			PrintDownFromTargetNode(n.left, k-2-rd)
		}
		return rd + 1
	}
	return -1
}

func main() {

	n := &node{data: 20}
	n.left = &node{data: 8}
	n.right = &node{data: 22}

	n.left.left = &node{data: 4}
	n.left.right = &node{data: 12}

	n.left.right.left = &node{data: 10}
	n.left.right.right = &node{data: 14}

	targetNode := 8
	k := 2
	target := findTargetNode(n, targetNode)

	printKthDistanceNodesFromGivenNode(n, target, k)

}

type node struct {
	data  int
	left  *node
	right *node
}
