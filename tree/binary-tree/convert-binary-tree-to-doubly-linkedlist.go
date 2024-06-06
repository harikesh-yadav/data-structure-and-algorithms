package main

import (
	"fmt"
)

/**

**/

func ConvertToDLL(root *node) *node {
	if root == nil {
		return nil
	}
	var head, prev *node
	var convert func(root *node)

	convert = func(root *node) {
		if root == nil {
			return
		}
		convert(root.left)

		if prev == nil {
			head = root
		} else {
			root.left = prev
			prev.right = root
		}
		prev = root
		convert(root.right)

	}
	convert(root)
	return head
}

func main() {
	bt := &binaryTree{}
	bt.insert(100)
	bt.insert(800)
	bt.insert(500)
	bt.insert(400)
	bt.insert(30)
	bt.insert(40)
	bt.insert(600)
	bt.insert(525)
	bt.insert(700)

	head := ConvertToDLL(bt.root)

	current := head
	for current != nil {
		fmt.Print(current.data, " ")
		current = current.right
	}
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
