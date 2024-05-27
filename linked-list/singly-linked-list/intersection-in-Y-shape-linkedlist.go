package main

import "fmt"

func lenght(ll *node) int {
	if ll == nil {
		return 0
	}
	size := 0
	current := ll
	for current != nil {
		size++
		current = current.next
	}
	return size
}

func findIntersectionPoint(l1, l2 *node) {
	if l1 == nil || l2 == nil {
		return
	}
	n := lenght(l1)
	m := lenght(l2)
	diff := 0
	var ptr1, ptr2 *node

	if n > m {
		ptr1 = l1
		ptr2 = l2
		diff = n - m
	} else {
		ptr1 = l2
		ptr1 = l1
		diff = m - n
	}

	for ptr1 != nil && diff != 0 {
		ptr1 = ptr1.next
		diff--
	}

	for ptr1 != nil && ptr2 != nil {
		if ptr1.data == ptr2.data {
			break
		}
		ptr1 = ptr1.next
		ptr2 = ptr2.next
	}
	fmt.Println("Intersection point is ", ptr1.data)

}

func main() {
	l1 := &node{data: 3}
	l1.next = &node{data: 6}
	l1.next.next = &node{data: 9}
	l1.next.next.next = &node{data: 15}
	l1.next.next.next.next = &node{data: 30}

	l2 := &node{data: 10}
	l2.next = &node{data: 15}
	l2.next.next = &node{data: 30}

	findIntersectionPoint(l1, l2)

}

type node struct {
	data int
	next *node
}
