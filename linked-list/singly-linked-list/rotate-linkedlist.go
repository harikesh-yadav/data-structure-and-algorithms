package main

import "fmt"

func rotateLeft(ll *node, k int) {
	if ll == nil {
		return
	}
	current := ll
	head2 := current
	for current != nil && k > 1 {
		current = current.next
		k--
	}
	if k == 1 && current != nil {
		ll = current.next
		current.next = nil
	}
	current = ll
	for current.next != nil {
		current = current.next
	}
	current.next = head2

	Print(ll)
}

func main() {

	l1 := &node{data: 2}
	l1.next = &node{data: 4}
	l1.next.next = &node{data: 7}
	l1.next.next.next = &node{data: 8}
	l1.next.next.next.next = &node{data: 9}

	rotateLeft(l1, 3)
}

func Print(ll *node) {
	if ll == nil {
		return
	}
	current := ll
	for current != nil {
		fmt.Print(current.data, ", ")
		current = current.next
	}
}

type node struct {
	data int
	next *node
}
