package main

import "fmt"

type mystack struct {
	arr []int
}

// push

func (s *mystack) push(item int) []int {
	s.arr = append(s.arr, item)
	return s.arr
}

// push
func (s *mystack) pop() int {
	item := s.arr[len(s.arr)-1]
	s.arr = s.arr[0 : len(s.arr)-1]
	return item
}

func (s *mystack) printstack() {
	fmt.Printf("%v", s.arr)
}

func IsValidParenteses(s string) bool {
	if s == "" {
		fmt.Println("invalid imput")
		return false
	}

	var stack []rune
	var maping = map[rune]rune{'}': '{', ')': '(', ']': '['}
	for _, c := range s {
		var item rune
		if c == ')' || c == ']' || c == '}' {
			if len(stack) > 0 {
				item = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			} else {
				item = '#'
			}
			if maping[c] != item {
				return false
			}
		} else {
			stack = append(stack, c)
		}
	}

	return len(stack) == 0
}
