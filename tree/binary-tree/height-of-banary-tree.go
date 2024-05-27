package main

import "fmt"

func heightBT(n *node) int {
	if n == nil {
		return 0
	}
	lh := heightBT(n.left)
	rh := heightBT(n.right)
	if lh > rh {
		return 1 + lh
	}
	return 1 + rh
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
	h := heightBT(bt.root)
	fmt.Println("height of Binary tree:=", h)

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
