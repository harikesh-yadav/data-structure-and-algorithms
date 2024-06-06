package main

import (
	"fmt"
	"math"
)

/*****
https://www.geeksforgeeks.org/a-program-to-check-if-a-binary-tree-is-bst-or-not/

Time Complexity: O(n), where n is the number of node in BST.
Auxiliary Space: O(1), if Function Call Stack size is not considered, otherwise O(H) where H is the height of the tree
******/

func checkBST(n *node, min, max int) bool {
	if n == nil {
		return true
	}

	if n.data < min || n.data > max {
		return false
	}
	return checkBST(n.left, min, n.data) && checkBST(n.right, n.data, max)
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

	ok := checkBST(bst, math.MinInt, math.MaxInt)

	fmt.Println(ok)
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
