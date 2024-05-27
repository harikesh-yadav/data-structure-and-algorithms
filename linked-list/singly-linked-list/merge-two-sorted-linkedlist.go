package main

import (
	"fmt"
)

func mergeTwoSortedLinkeslist(ll1, ll2 *singlyLinkedList) {
	if ll1.head == nil {
		ll2.Print()
		return
	}
	if ll2.head == nil {
		ll1.Print()
		return
	}

	ll := &singlyLinkedList{}

	current1 := ll1.head
	current2 := ll2.head

	for current1 != nil && current2 != nil {
		if current1.data == current2.data {
			ll.InsertAtEnd(current1.data)
			ll.InsertAtEnd(current2.data)
			current1 = current1.next
			current2 = current2.next
		} else if current1.data < current2.data {
			ll.InsertAtEnd(current1.data)
			current1 = current1.next
		} else {
			ll.InsertAtEnd(current2.data)
			current2 = current2.next
		}
	}

	if current1 != nil {
		for current1 != nil {
			ll.InsertAtEnd(current1.data)
			current1 = current1.next
		}
	}
	if current2 != nil {
		for current2 != nil {
			ll.InsertAtEnd(current2.data)
			current2 = current2.next
		}
	}
	ll.Print()
}

func main() {
	sl1 := &singlyLinkedList{}
	sl1.InsertAtEnd(1)
	sl1.InsertAtEnd(2)
	sl1.InsertAtEnd(3)
	sl1.InsertAtEnd(5)
	sl1.InsertAtEnd(7)
	sl1.InsertAtEnd(9)
	//sl1.Print()
	fmt.Println()
	sl2 := &singlyLinkedList{}
	sl2.InsertAtEnd(2)
	sl2.InsertAtEnd(4)
	sl2.InsertAtEnd(6)
	sl2.InsertAtEnd(8)
	sl2.InsertAtEnd(10)
	sl2.Print()
	fmt.Println()
	mergeTwoSortedLinkeslist(sl1, sl2)

}

type node struct {
	data int
	next *node
}

type singlyLinkedList struct {
	head *node
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
