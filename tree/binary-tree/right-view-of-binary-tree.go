package main

import (
	"fmt"
)

/**
Problem:

time complexity:  O(n)
space complexity :O(n)  // n is number of node present on same level
**/

func rightViewOnBinaryTree(n *node, hash *map[int]int, level int) {
	if n == nil {
		return
	}
	if _, ok := (*hash)[level]; !ok {
		(*hash)[level] = n.data
	}
	rightViewOnBinaryTree(n.right, hash, level+1)
	rightViewOnBinaryTree(n.left, hash, level+1)

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
	hashmap := make(map[int]int)
	initialLevel := 0
	rightViewOnBinaryTree(bt.root, &hashmap, initialLevel)

	for i := 0; i < len(hashmap); i++ {
		fmt.Println(hashmap[i])
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
