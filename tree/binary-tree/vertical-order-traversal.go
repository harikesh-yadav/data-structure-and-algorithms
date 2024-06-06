package main

import (
	"fmt"
)

/*****

time complexity o(n)
space complexity o(n)

*****/

func verticalOrderTraversal(n *node, level int) {
	if n == nil {
		return
	}

	if len(set[level]) > 0 {
		set[level] = append(set[level], n.data)
	} else {
		set[level] = []int{n.data}
	}

	verticalOrderTraversal(n.left, level-1)
	verticalOrderTraversal(n.right, level+1)

}

var set map[int][]int

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

	set = make(map[int][]int)

	verticalOrderTraversal(bst, 0)

	for k, arr := range set {
		fmt.Print(k, ": ")
		for _, v := range arr {
			fmt.Print(v, " ")
		}
		fmt.Println()
	}

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
