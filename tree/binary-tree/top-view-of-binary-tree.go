package main

import (
	"fmt"
)

/**

IN TOP VIEW WE HAVE TO DO LEVEL ORDER TRAVERSAL AND
STORE HORIZONTAL DISTANCE AND NODE PAIR IN QUEUE. (pair)

Problem:  https://www.youtube.com/watch?v=wTloJwckQTU&list=PLUcsbZa0qzu0g_LceXxWu6ICc2o8SiAhJ&index=7

time complexity:  O(n)
space complexity :O(n)  // n is number of node present on same level
**/

type pair struct {
	hd    int
	tNode *node
}

func topViewOnBinaryTree(n *node, hash map[int]int, hd int) map[int]int {
	if n == nil {
		return nil
	}

	queue := make([]pair, 0)
	queue = append(queue, pair{hd, n})
	for len(queue) > 0 {
		front := queue[0]
		if _, ok := hash[front.hd]; !ok {
			hash[front.hd] = front.tNode.data
		}
		queue = queue[1:]
		if front.tNode.left != nil {
			lhd := front.hd - 1
			queue = append(queue, pair{lhd, front.tNode.left})
		}
		if front.tNode.right != nil {
			rhd := front.hd + 1
			queue = append(queue, pair{rhd, front.tNode.right})
		}
	}
	return hash

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

	topViewOnBinaryTree(bt.root, hashmap, initialLevel)

	for _, v := range hashmap {
		fmt.Println(v)
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
