package main

import (
	"fmt"
	"math"
)

func maximum(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

func maxNode(n *node) int {
	if n == nil {
		return math.MinInt
	}
	l := maxNode(n.left)
	r := maxNode(n.right)

	m1 := maximum(l, r)
	m2 := maximum(m1, n.data)
	return m2
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
	//print(bt.root, 1, "ROOT")
	//max := math.MinInt
	max := maxNode(bt.root)
	fmt.Println("MAX node of Binary tree:=", max)

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
