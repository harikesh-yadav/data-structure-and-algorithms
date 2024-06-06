package main

import (
	"fmt"
	"math"
)

/*****
https://www.geeksforgeeks.org/floor-and-ceil-from-a-bst/

Time Complexity: O(log N), where N is the number of node in BST.
Auxiliary Space:

******/

func findCeilValue(n *node, key int) int {
	if n == nil {
		return math.MaxInt
	}

	result := math.MaxInt
	for n != nil {
		if n.data == key {
			return n.data
		}

		if n.data < key {
			n = n.right
		} else {
			result = n.data
			n = n.left
		}
	}
	return result
}

func main() {

	bst := &node{data: 30}
	bst.left = &node{data: 28}
	bst.right = &node{data: 40}

	bst.left.left = &node{data: 20}

	bst.left.left.left = &node{data: 10}
	bst.left.left.right = &node{data: 25}

	bst.left.left.left.left = &node{data: 5}

	bst.left.left.right.left = &node{data: 24}
	bst.left.left.right.right = &node{data: 26}

	f := findCeilValue(bst, 22)

	fmt.Println(f)

	printBST(bst)
}

type node struct {
	data  int
	left  *node
	right *node
}

func printBST(n *node) {
	if n == nil {
		return
	}

	printBST(n.left)
	fmt.Print(n.data, " ")
	printBST(n.right)
}
