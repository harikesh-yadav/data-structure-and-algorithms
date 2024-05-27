package main

import (
	"fmt"
	"log"
)

type node struct {
	data int
	next *node
}

type singlyLinkedList struct {
	head *node
}

/***********************INSERT **********************************/

func (ll *singlyLinkedList) InsertAtBegning(n int) {
	node := &node{data: n}
	second := ll.head
	ll.head = node
	ll.head.next = second
}

func (ll *singlyLinkedList) InsertAfter(n, after int) {
	if ll.head == nil {
		return
	}
	current := ll.head
	for current.next != nil && current.data != after {
		current = current.next
	}
	if current.data == after {
		node := &node{data: n}
		next := current.next
		current.next = node
		current.next.next = next
		return
	}
	log.Printf("%d node not found!\n", after)
}

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
	// TO MAKE CIRCULAR
	// if n == 5 {
	// 	node.next = current
	// }
	current.next = node

}

/***********************DELETE **********************************/

func (ll *singlyLinkedList) DeleteAtBigening() {
	if ll.head == nil {
		fmt.Println("Empty Linkedlist")
		return
	}
	second := ll.head.next
	ll.head = second
}

func (ll *singlyLinkedList) DeleteAfter(after int) {
	if ll.head == nil {
		fmt.Println("Empty Linkedlist")
		return
	}
	current := ll.head
	for current.next != nil && current.data != after {
		current = current.next
	}
	if current.data == after {
		if current.next != nil {
			next := current.next.next
			current.next = next
		} else {
			current.next = nil
		}
	} else {
		log.Printf("%d node not found!\n", after)
	}

}

func (ll *singlyLinkedList) DeleteAtEnd() {
	if ll.head == nil {
		fmt.Println("Empty Linkedlist")
		return
	}

	current := ll.head

	for current.next != nil && current.next.next != nil {
		current = current.next
	}
	current.next = nil
}

/***********************TRAVERSE OR PRINT **********************************/

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

func (ll *singlyLinkedList) Search(n int) {
	if ll.head == nil {
		return
	}

	current := ll.head
	for current != nil {
		if current.data == n {
			fmt.Println(n)
			return
		}
		current = current.next
	}
}

func (ll *singlyLinkedList) Reverse() {
	if ll.head == nil {
		return
	}
	current := ll.head
	var prev *node
	var next *node
	for current != nil {
		next = current.next
		current.next = prev
		prev = current
		current = next
	}
	ll.head = prev
}

func (ll *singlyLinkedList) IsCircular() {
	if ll.head == nil {
		fmt.Print("1: LinkedlIst is not Circular ")
		return

	}
	slow := ll.head
	fast := ll.head.next
	for fast != nil && fast.next != nil {
		if slow == fast {
			fmt.Print("LinkedlIst is Circular ")
			return
		}
		slow = slow.next
		fast = fast.next.next
	}
	fmt.Print("2: LinkedlIst is not Circular ")
}

func main() {
	sl := &singlyLinkedList{}
	sl.InsertAtEnd(1)
	sl.InsertAtEnd(2)
	sl.InsertAtEnd(3)
	sl.InsertAtEnd(4)
	sl.InsertAtEnd(5)
	//sl.Reverse()

	sl.IsCircular()
	//sl.Print()
}
