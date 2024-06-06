package main

import (
	"fmt"
)

/**
Time Complexity: O(N) where N is the total number of nodes. Because it traverses all the nodes at least once.
Auxiliary Space: O(N)
**/

func inOrderIterative(n *node) {
	stack := []*node{}
	current := n
	for current != nil || len(stack) > 0 {
		for current != nil {
			stack = append(stack, current)
			current = current.left
		}
		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		fmt.Print(current.data, " ")
		current = current.right
	}

}

func preOrderIterative(n *node) {
	stack := make([]*node, 0)
	stack = append(stack, n)
	var current *node

	for len(stack) > 0 {
		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		fmt.Print(current.data, " ")

		if current.right != nil {
			stack = append(stack, current.right)
		}
		if current.left != nil {
			stack = append(stack, current.left)
		}
	}
}

// Problem:  https://www.youtube.com/watch?v=NzIGLLwZBS8

func postOrderIterative(n *node) {
	stack := []*node{}
	current := n

	for current != nil || len(stack) > 0 {
		for current != nil {
			stack = append(stack, current)
			current = current.left
		}
		temp := stack[len(stack)-1].right

		if temp == nil {
			temp = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			fmt.Print(temp.data, " ")
			for len(stack) > 0 && temp == stack[len(stack)-1].right {
				temp = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				fmt.Print(temp.data, " ")
			}

		} else {
			current = temp
		}
	}
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

	//inOrderIterative(bt.root)
	// preOrderIterative(bt.root)
	postOrderIterative(bt.root)

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
