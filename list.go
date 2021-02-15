package main

import "fmt"

type node struct {
	data int
	next *node
}
type mylist struct {
	head  *node
	count int
}

func intitList() *mylist {
	return &mylist{}
}

func (s *mylist) insert(value int) {
	newdata := &node{data: value, next: nil}

	if s.head == nil {
		s.head = newdata
		s.count++
		return
	}

	temp := s.head
	for ; temp.next != nil; temp = temp.next {
	}
	temp.next = newdata
	s.count++
	return
}

func (head *node) display() {
	if head == nil {
		fmt.Println(" List is empty ")
		return
	}

	for temp := head; temp != nil; temp = temp.next {
		fmt.Printf("%d ==>", temp.data)
	}
	fmt.Println()
	return
}

func (s *mylist) reverseList() {
	if s.head == nil {
		fmt.Println("Empty list")
		return
	}

	current := s.head
	var prev *node
	var next *node
	for current != nil {
		next = current.next
		current.next = prev
		prev = current
		current = next
	}
	s.head = prev
}

func reverseListRec(head *node) *node {
	if head == nil || head.next == nil {
		return head
	}
	var rest *node
	rest = reverseListRec(head.next)
	fmt.Println(head.next.next)
	fmt.Println(head.next)
	fmt.Println("Rest value", rest)
	head.next.next = head
	head.next = nil

	return rest
}

func mergeSortedList(l1 *node, l2 *node) *node {
	head := &node{}
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	if l1.data <= l2.data {
		head = l1
		head.next = mergeSortedList(l1.next, l2)
	} else {
		head = l2
		head.next = mergeSortedList(l1, l2.next)
	}
	return head
}
