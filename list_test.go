package main

import (
	"fmt"
	"testing"
)

func TestListAdd(t *testing.T) {
	head := intitList()
	head.insert(10)
	head.insert(20)
	head.insert(30)
	// head.insert(40)
	// head.insert(50)
	// head.insert(60)
	// head.insert(70)
	// head.insert(80)
	head.head.display()

	// fmt.Println("Reverse List count ", head.count)
	// head.reverseList()
	// head.head.display()

	fmt.Println("Reverse List Recurrsive 0-count ", head.count)
	fmt.Println(head.head)
	list := reverseListRec(head.head)
	fmt.Println(list)
	list.display()
}

func TestMergeSortedList(t *testing.T) {
	head1 := &node{data: 10, next: nil}
	head1.next = &node{data: 20, next: nil}
	head1.next.next = &node{data: 60, next: nil}
	head2 := &node{data: 30, next: nil}
	head2.next = &node{data: 50, next: nil}
	head1.next.next = &node{data: 70, next: nil}
	result := mergeSortedList(head1, head2)
	result.display()
}
