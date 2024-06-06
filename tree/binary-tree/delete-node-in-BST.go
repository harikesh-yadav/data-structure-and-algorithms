package main

import (
	"fmt"
)

/*****
https://www.geeksforgeeks.org/deletion-in-binary-search-tree/

Time Complexity: O(h), where h is the height of the BST.
Auxiliary Space: O(n)

******/

func min(n *node) int {
	min := n.data
	for n.left != nil {
		min = n.left.data
		n = n.left
	}
	return min
}

func deleteNode(n *node, key int) *node {
	if n == nil {
		return nil
	}

	if key < n.data {
		n.left = deleteNode(n.left, key)
	} else if key > n.data {
		n.right = deleteNode(n.right, key)
	} else {
		if n.left == nil {
			return n.right
		} else if n.right == nil {
			return n.left
		} else {
			// if left and right both not nil
			n.data = min(n.right)
			n.right = deleteNode(n.right, n.data)

		}
	}
	return n
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

	printBST(bst)
	fmt.Println()

	deleteNode(bst, 20)

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
