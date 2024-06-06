package main

import (
	"fmt"
)

/*****
https://www.geeksforgeeks.org/find-a-pair-with-given-sum-in-bst/

Time Complexity: O(N), where N is the number of node in BST.
Auxiliary Space: O(N)

******/

func pairOFSumInBST(n *node, sum int) bool {
	if n == nil {
		return false
	}

	ok := pairOFSumInBST(n.left, sum)
	if ok {
		return true
	}

	if set[sum-n.data] {
		return true
	} else {
		set[n.data] = true
	}
	return pairOFSumInBST(n.right, sum)
}

var set map[int]bool

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
	set = make(map[int]bool)

	ok := pairOFSumInBST(bst, 50)

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
