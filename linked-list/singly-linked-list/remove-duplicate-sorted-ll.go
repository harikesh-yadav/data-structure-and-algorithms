package main

import (
	"fmt"
)

func removeDuplicate(ll *singlyLinkedList) {

	if ll.head == nil {
		return
	}

	prev := ll.head
	current := ll.head.next
	for current != nil && current.next != nil {
		if prev.data == current.data {
			current = current.next
		} else {
			prev.next = current
			prev = prev.next
			current = current.next
		}
	}
	if current != nil && prev.data != current.data {
		prev.next = current
	} else {
		prev.next = nil
	}

}

func main() {
	ll := &singlyLinkedList{}
	ll.InsertAtEnd(1)
	ll.InsertAtEnd(1)
	ll.InsertAtEnd(2)
	ll.InsertAtEnd(2)
	ll.InsertAtEnd(2)
	ll.InsertAtEnd(3)
	ll.InsertAtEnd(3)
	ll.InsertAtEnd(3)

	removeDuplicate(ll)
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
