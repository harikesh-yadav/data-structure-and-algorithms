package main

import (
	"fmt"
)

/**
Problem:   https://www.geeksforgeeks.org/delete-middle-of-linked-list/
*/

func findAndDeleteMiddleElement(ll *singlyLinkedList) {
	if ll.head == nil {
		return
	}
	slow := ll.head
	fast := ll.head
	prev := ll.head

	for fast != nil && fast.next != nil {
		prev = slow
		slow = slow.next
		fast = fast.next.next
	}

	prev.next = slow.next

}

func main() {
	ll := &singlyLinkedList{}
	ll.InsertAtEnd(1)
	ll.InsertAtEnd(2)
	ll.InsertAtEnd(3)
	ll.InsertAtEnd(4)
	ll.InsertAtEnd(5)
	ll.InsertAtEnd(6)

	findAndDeleteMiddleElement(ll)
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
