package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//Definition for a binary tree node..
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func ArrayToBST(input []int) *TreeNode {
	if len(input) == 0 {
		return nil
	}
	root := helperTree(input)
	return root
}

func helperTree(values []int) *TreeNode {
	if len(values) == 0 {
		return nil
	}
	mid := len(values) / 2
	return &TreeNode{
		Val:   values[mid],
		Left:  helperTree(values[:mid]),
		Right: helperTree(values[mid+1:]),
	}
}

type Node struct {
	val         interface{}
	left, right *Node
}

// PreOrder converts a tree rooted at given node into a linked list
// ... that represents the pre-order traversal of the original tree.

type Stack []*Node

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(str *Node) {
	*s = append(*s, str)
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (*Node, bool) {
	element := &Node{}
	if s.IsEmpty() {
		return element, false
	} else {
		index := len(*s) - 1
		element = (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}

// Get top element from stack

func (s *Stack) Top() (*Node, bool) {
	element := &Node{}
	if s.IsEmpty() {
		return element, false
	} else {
		index := len(*s) - 1  // Get the index of the top most element.
		element = (*s)[index] // Index into the slice and obtain the element.
		return element, true
	}
}

// function to convert tree to list in preorder using stack
func PreOrderFlatten(root *Node) {

	if root == nil {
		return
	}

	var mystack Stack

	mystack.Push(root)

	for !mystack.IsEmpty() {

		temp, _ := mystack.Pop()
		if temp.right != nil {
			mystack.Push(temp.right)
		}
		if temp.left != nil {
			mystack.Push(temp.left)
		}

		if !mystack.IsEmpty() {
			temp.right, _ = mystack.Top()
		}

		temp.left = nil
	}

	PrintTree(root)
}

// Print the linked list

func PrintTree(root *Node) {
	for temp := root; temp != nil; temp = temp.right {
		fmt.Printf("%d ", temp.val)
	}
}

/*
				1
		2				3
  4		5		6		7

  Result [1,2,4,5,3,6,7]
*/

func PrintPreorder(root *Node) {
	if root == nil {
		return
	}
	fmt.Printf(" %d ", root.val)
	PrintPreorder(root.left)
	PrintPreorder(root.right)
}

// Result [1,2,4,5,3,6,7]
func Inorderder(root *Node) {
	if root == nil {
		return
	}
	PrintPreorder(root.left)
	fmt.Printf(" %d ", root.val)
	PrintPreorder(root.right)
}
