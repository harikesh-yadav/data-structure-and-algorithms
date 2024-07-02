package main

import (
	"fmt"
)

/**
Problem:   https://www.geeksforgeeks.org/detect-and-remove-loop-in-a-linked-list/

time complexity : O(n)
space complexity: O(n)

*/

func isCircular(ll *singlyLinkedList) bool {
	if ll.head == nil {
		return false
	}
	slow := ll.head
	fast := ll.head
	for fast != nil && fast.next != nil {
		if slow == fast {
			return true
		}
		slow = slow.next
		fast = fast.next.next
	}
	return false

}

func removeLoop(ll *singlyLinkedList) {
	if ll.head == nil {
		return
	}

	if !isCircular(ll) {
		fmt.Println("Not Circular ")
		return
	} else {
		fmt.Println(" Circular ")
	}
	slow := ll.head
	fast := ll.head.next

	for fast != nil && fast.next != nil {

		if slow == fast {
			break
		}

		slow = slow.next
		fast = fast.next.next
	}
	slow = ll.head
	for fast.next != slow {
		slow = slow.next
		fast = fast.next
	}
	fast.next = nil

}

func main() {
	ll := &singlyLinkedList{}
	ll.InsertAtEnd(1)
	ll.InsertAtEnd(2)
	ll.InsertAtEnd(3)
	ll.InsertAtEnd(4)
	ll.InsertAtEnd(5)

	removeLoop(ll)
	ll.Print()

}

type node struct {
	data int
	next *node
}

type singlyLinkedList struct {
	head *node
}

/***********************INSERT **********************************/

func (ll *singlyLinkedList) InsertAtEnd(n int) {
	node := &node{data: n}
	if ll.head == nil {
		ll.head = node
		return
	}
	current := ll.head
	for current.next != nil {
		current = current.next
	}
	// TO MAKE LINKED LIST CIRCULAR
	if n == 5 {
		node.next = ll.head
	}

	current.next = node
}

func (ll *singlyLinkedList) Print() {
	if ll.head == nil {
		fmt.Println("Linkedlist is empty.")
		return
	}
	current := ll.head
	for current != nil {
		fmt.Print(current.data, ", ")
		current = current.next
	}
}
