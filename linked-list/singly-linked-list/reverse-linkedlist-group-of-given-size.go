package main

import (
	"fmt"
)

func reverseInGroup(ll *node, k int) {
	if ll == nil || k <= 1 {
		return
	}
	current := ll
	var head, prev, next *node
	count := 1
	temp1 := current
	var temp2 *node
	for current != nil {
		for current.next != nil && count < k {
			next = current.next
			current.next = prev
			prev = current
			current = next
			count++
		}
		if current != nil {
			next = current.next
		}

		current.next = prev
		prev = nil

		if head == nil {
			head = current
		} else {
			temp1.next = current
			temp1 = temp2
		}
		current = next
		temp2 = current
		count = 1
	}

	Print(head)

}

func main() {

	// k = 4
	l1 := &node{data: 1}
	l1.next = &node{data: 2}
	l1.next.next = &node{data: 3}
	l1.next.next.next = &node{data: 4}
	l1.next.next.next.next = &node{data: 5}
	l1.next.next.next.next.next = &node{data: 6}
	l1.next.next.next.next.next.next = &node{data: 7}
	l1.next.next.next.next.next.next.next = &node{data: 8}
	l1.next.next.next.next.next.next.next.next = &node{data: 9}

	reverseInGroup(l1, 3)
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
