package main

import (
	"fmt"
)

func mergeSortedLL(l1, l2 *nodeWithBottom) *nodeWithBottom {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	temp1 := l1
	temp2 := l2
	head := &nodeWithBottom{}
	for temp1 != nil && temp2 != nil {

		if temp1.data < temp2.data {
			head = temp1.InsertAtEnd(temp1.data)
			temp1 = temp1.bottom
		} else if temp1.data > temp2.data {
			head = temp2.InsertAtEnd(temp2.data)
			temp2 = temp2.bottom
		}
	}

	for temp1 != nil {
		head = temp1.InsertAtEnd(temp1.data)
		temp1 = temp1.bottom

	}
	for temp2 != nil {
		head = temp2.InsertAtEnd(temp2.data)
		temp2 = temp2.bottom
	}

	return head
}

func flatten(ll *nodeWithBottom) *nodeWithBottom {
	if ll == nil {
		return nil
	}

	var temp1, temp2 *nodeWithBottom
	current := ll
	for current != nil && current.next != nil {
		if temp1 == nil {
			temp1 = current
		}
		temp2 = current.next

		temp1 = mergeSortedLL(temp1, temp2)

		current = current.next
	}

	return temp1
}

func main() {

	//Output: 5->7->8->10->19->20->22->28->30->35->40->45->50
	ll := &nodeWithBottom{data: 5}
	ll.bottom = &nodeWithBottom{data: 7}
	ll.bottom.bottom = &nodeWithBottom{data: 8}
	ll.bottom.bottom.bottom = &nodeWithBottom{data: 30}

	ll.next = &nodeWithBottom{data: 10}
	ll.next.bottom = &nodeWithBottom{data: 20}

	ll.next.next = &nodeWithBottom{data: 19}
	ll.next.next.bottom = &nodeWithBottom{data: 22}
	ll.next.next.bottom.bottom = &nodeWithBottom{data: 50}

	ll.next.next.next = &nodeWithBottom{data: 28}
	ll.next.next.next.bottom = &nodeWithBottom{data: 35}
	ll.next.next.next.bottom.bottom = &nodeWithBottom{data: 40}
	ll.next.next.next.bottom.bottom.bottom = &nodeWithBottom{data: 45}

	printNextBottomNode(flatten(ll))
}

type nodeWithBottom struct {
	data   int
	next   *nodeWithBottom
	bottom *nodeWithBottom
}

func printNextBottomNode(ll *nodeWithBottom) {
	if ll == nil {
		return
	}

	current := ll
	for current != nil {
		fmt.Print(current.data, ", ")
		var bottom *nodeWithBottom
		if current.bottom != nil {
			bottom = current.bottom
		}
		for bottom != nil {
			fmt.Print(bottom.data, ", ")
			bottom = bottom.bottom
		}
		current = current.next
	}
}

func (ll *nodeWithBottom) InsertAtEnd(n int) *nodeWithBottom {
	node := &nodeWithBottom{data: n}
	if ll == nil {
		ll = node
		return ll
	}
	current := ll
	for current.next != nil {
		current = current.next
	}
	current.next = node
	return ll
}
