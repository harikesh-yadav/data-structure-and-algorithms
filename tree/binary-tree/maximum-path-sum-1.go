package main

import (
	"fmt"
	"math"
)

/**

Problem: for Any node to any node in Binary Tree

https://www.geeksforgeeks.org/find-maximum-path-sum-in-a-binary-tree/

Time Complexity: O(N)  where N is the number of nodes in the Binary Tree
Auxiliary Space: O(N) due to recursive calls.
**/

func height(n *node, result *int) int {
	if n == nil {
		return 0
	}
	lh := height(n.left, result)
	rh := height(n.right, result)

	max1 := maximum(lh, rh) + n.data
	max2 := maximum(max1, n.data)
	ans := maximum(max2, lh+rh+n.data)
	*result = maximum(ans, *result)
	return max2

}

func maximum(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

func main() {
	bt := &binaryTree{}
	bt.insert(10)
	bt.insert(2)
	bt.insert(10)
	bt.insert(20)
	bt.insert(1)
	bt.insert(-25)
	bt.insert(3)
	bt.insert(4)
	print(bt.root, 1, "ROOT")

	result := math.MinInt
	height(bt.root, &result)
	fmt.Println("Maximum Path Sum | From any node to any node= ", result)

}

type node struct {
	data  int
	left  *node
	right *node
}

func (n *node) insert(d int) {
	if n == nil {
		return
	}
	if n.data > d {
		if n.left == nil {
			n.left = &node{data: d}
		} else {
			n.left.insert(d)
		}
	} else {
		if n.right == nil {
			n.right = &node{data: d}
		} else {
			n.right.insert(d)
		}
	}
}

type binaryTree struct {
	root *node
}

func (bt *binaryTree) insert(d int) {
	if bt.root == nil {
		bt.root = &node{data: d}
		return
	}
	bt.root.insert(d)
}

func print(n *node, sp int, ntype string) {
	if n == nil {
		return
	}

	for i := 0; i < sp; i++ {
		fmt.Print("  ")
	}
	fmt.Print(ntype + ":")
	fmt.Println(n.data)

	print(n.left, sp+1, "L")
	print(n.right, sp+1, "R")

}
