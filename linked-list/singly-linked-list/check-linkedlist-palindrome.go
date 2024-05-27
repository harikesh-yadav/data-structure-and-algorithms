package main

import (
	"fmt"
)

/**
Problem:
*/

func findMiddleElement(ll *singlyLinkedList) *node {

	if ll.head == nil {
		return nil
	}

	slow := ll.head
	fast := ll.head.next
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}

	return slow
}

func reverse(ll *node) *node {
	if ll == nil {
		return nil
	}

	var prev, next *node
	current := ll

	for current.next != nil {
		next = current.next
		current.next = prev
		prev = current
		current = next
	}
	return current
}

func checkPolidrome(ll *singlyLinkedList) bool {

	if ll.head == nil {
		return false
	}
	middle := findMiddleElement(ll)
	back := reverse(middle)
	front := ll.head
	for front != nil {
		front = front.next
	}

	for front != nil && back != nil {
		if front.data != back.data {
			return false
		}
		front = front.next
		back = back.next
	}
	return true
}

func main() {
	ll := &singlyLinkedList{}
	ll.InsertAtEnd(1)
	ll.InsertAtEnd(2)
	ll.InsertAtEnd(1)
	// ll.InsertAtEnd(4)
	// ll.InsertAtEnd(5)

	fmt.Println(checkPolidrome(ll))

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
