package main

import "fmt"

/***
Problem:  https://www.geeksforgeeks.org/problems/add-two-numbers-represented-by-linked-lists/1?page=1&category=Linked%20List&company=Microsoft&sortBy=submissions
**/

func addTwoNumsReperesentedByLinkedlist(l1, l2 *singlyLinkedList) {
	if l1 == nil || l2 == nil {
		return
	}

	num1 := 0
	num2 := 0
	first := l1.head
	second := l2.head

	for first != nil && second != nil {
		num1 = num1*10 + first.data
		num2 = num2*10 + second.data
		first = first.next
		second = second.next
	}

	if first != nil {
		for first != nil {
			num1 = num1*10 + first.data
			first = first.next
		}
	}
	if second != nil {
		for second != nil {
			num2 = num2*10 + second.data
			second = second.next
		}
	}

	sum := num1 + num2
	stack := make([]int, 0)
	for sum != 0 {
		r := sum % 10
		stack = append([]int{r}, stack...)
		sum = sum / 10
	}

	ll := &singlyLinkedList{}
	for _, v := range stack {
		ll.InsertAtEnd(v)
	}
	ll.Print()

}

func main() {
	ll := &singlyLinkedList{}
	ll.InsertAtEnd(0)
	ll.InsertAtEnd(0)
	ll.InsertAtEnd(6)
	ll.InsertAtEnd(3)

	l2 := &singlyLinkedList{}
	l2.InsertAtEnd(0)
	l2.InsertAtEnd(7)

	addTwoNumsReperesentedByLinkedlist(ll, l2)

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
