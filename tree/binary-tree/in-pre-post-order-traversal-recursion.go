package main

import (
	"fmt"
)

/**
Time Complexity: O(N) where N is the total number of nodes. Because it traverses all the nodes at least once.
Auxiliary Space: O(1) if no recursion stack space is considered. Otherwise, O(h) where h is the height of the tree
**/

func inOrder(n *node) {
	if n == nil {
		return
	}
	inOrder(n.left)
	fmt.Println(n.data)
	inOrder(n.right)
}

func preOrder(n *node) {
	if n == nil {
		return
	}
	fmt.Println(n.data)
	preOrder(n.left)
	preOrder(n.right)
}

func postOrder(n *node) {
	if n == nil {
		return
	}
	postOrder(n.left)
	postOrder(n.right)
	fmt.Println(n.data)
}

func main() {
	bt := &binaryTree{}
	bt.insert(100)
	bt.insert(20)
	bt.insert(500)
	bt.insert(400)
	bt.insert(30)
	bt.insert(40)
	bt.insert(600)
	bt.insert(525)
	bt.insert(700)
	//print(bt.root, 1, "ROOT")
	//inOrder(bt.root)
	preOrder(bt.root)
	//postOrder(bt.root)

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
