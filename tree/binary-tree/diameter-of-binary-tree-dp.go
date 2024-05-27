package main

import (
	"fmt"
	"math"
)

/**
https://www.geeksforgeeks.org/diameter-of-a-binary-tree/
Time Complexity: O(N)
Auxiliary Space: O(N) due to recursive calls.
**/

func height(n *node, result *int) int {
	if n == nil {
		return 0
	}
	lh := height(n.left, result)
	rh := height(n.right, result)
	mh := maximum(lh, rh) + 1
	ans := maximum(mh, lh+rh+1)
	*result = maximum(ans, *result)
	return mh

}

func maximum(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

func main() {
	bt := &binaryTree{}
	bt.insert(100)
	bt.insert(20)
	bt.insert(500)
	bt.insert(10)
	bt.insert(30)
	bt.insert(40)
	bt.insert(600)
	bt.insert(650)
	bt.insert(700)
	//print(bt.root, 1, "ROOT")

	result := math.MinInt
	height(bt.root, &result)
	fmt.Println("Diameter of Tree:= ", result)

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
