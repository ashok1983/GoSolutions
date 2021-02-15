package main

import (
	"fmt"
	"testing"
)

func printtree(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	printtree(root.Left)
	printtree(root.Right)
}

func TestArrayToBST(t *testing.T) {
	input := []int{-10, -3, 0, 5}
	root := ArrayToBST(input)
	printtree(root)
}

/*
				1
		2				3
	4		5		6		7
*/
func TestPreorderFlatten(t *testing.T) {
	root := &Node{val: 1, left: nil, right: nil}
	root.left = &Node{val: 2, left: nil, right: nil}
	root.right = &Node{val: 3, left: nil, right: nil}
	root.left.left = &Node{val: 4, left: nil, right: nil}
	root.left.right = &Node{val: 5, left: nil, right: nil}
	root.right.left = &Node{val: 6, left: nil, right: nil}
	root.right.right = &Node{val: 7, left: nil, right: nil}
	PrintPreorder(root)
	fmt.Println()
	PreOrderFlatten(root)
	PrintTree(root)
}
