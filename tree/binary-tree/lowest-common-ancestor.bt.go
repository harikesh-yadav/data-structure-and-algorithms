package main

import "fmt"



func lowestCommonAncestor(n *node, n1, n2 int) *node {
	if n == nil {
		return nil
	}
	if n1 == n.data || n2 == n.data {
		return n
	}
	l := lowestCommonAncestor(n.left, n1, n2)
	r := lowestCommonAncestor(n.right, n1, n2)

	if l != nil && r != nil {
		return n
	}
	if l != nil && r == nil {
		return l
	}
	if l == nil && r != nil {
		return r
	}
	return nil
}

func main() {

	n := &node{data: 1}
	n.left = &node{data: 2}
	n.right = &node{data: 3}

	n.left.left = &node{data: 4}
	n.left.right = &node{data: 5}

	n.right.left = &node{data: 6}
	n.right.right = &node{data: 7}
	n1 := 
	n2 := 7
	fmt.Printf("Common Ancestor for %d and %d = %d", n1, n2, lowestCommonAncestor(n, n1, n2).data)

}

type node struct {
	data  int
	left  *node
	right *node
}
