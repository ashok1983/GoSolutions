package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicateMap(*testing.T) {
	input := []int{1, 1, 2, 2, 3, 4, 4, 4, 5, 5, 6, 7, 7, 8, 9, 9}
	RemoveDuplicate(input, len(input))

	input = []int{1, 1, 2, 2, 3, 4, 4, 4, 5, 5, 6, 7, 7, 8, 9, 9}
	RemoveDuplicateSameArray(input, len(input))
}
func TestIsRoatation(t *testing.T) {
	s1 := "waterbottlewaterbottle"
	s2 := "erbottlewat"
	result := IsRotation(s1, s2)
	fmt.Println("result is ", result)
	assert.Equal(t, true, result)

	s1 = "waterbottlewaterbottle"
	s2 = "hello"
	result = IsRotation(s1, s2)
	fmt.Println("result is ", result)
	assert.Equal(t, false, result)
}

func TestIsUnique(t *testing.T) {
	s1 := "hello"
	assert.Equal(t, false, isUniqueChars(s1))
	s1 = "abcdefghijklmnop"
	assert.Equal(t, true, isUniqueChars(s1))
}

func TestBalancedString(t *testing.T) {
	s1 := "RLRRLLRLRL"
	assert.Equal(t, 4, balancedString(s1))
	s1 = "RLLLLRRRLR"
	assert.Equal(t, 3, balancedString(s1))
	s1 = "LLLLRRRR"
	assert.Equal(t, 1, balancedString(s1))
	s1 = "RL"
	balancedString(s1)
	assert.Equal(t, 1, balancedString(s1))
}

func TestIspermutation(t *testing.T) {
	assert.Equal(t, true, Ispermutation("dog", "god"))
	assert.Equal(t, false, Ispermutation("dog", "pod"))
	assert.Equal(t, false, Ispermutation("", "god"))
	assert.Equal(t, false, Ispermutation("", ""))
}

func TestIspermutation2(t *testing.T) {
	assert.Equal(t, true, Ispermutation2("dog", "god"))
	assert.Equal(t, false, Ispermutation2("dog", "pod"))
	assert.Equal(t, false, Ispermutation2("", "god"))
	assert.Equal(t, false, Ispermutation2("", ""))
}
func TestIspermutationOfPalindrome(t *testing.T) {
	assert.Equal(t, true, IspermutationOfPalindrome("tactcoapapa"))
	assert.Equal(t, true, IspermutationOfPalindrome("aabb"))

}

func TestGeneratePermutation(t *testing.T) {
	s1 := "ABC"
	input := []rune(s1)
	generatePermutation(input, 0, len(input)-1)
	fmt.Println(s1)
}

func TestGeneratePermutationByte(t *testing.T) {
	s1 := "ABC"
	input := []byte(s1)
	generatePermutationByte(input, 0, len(input)-1)

}
func TestFibonacci(t *testing.T) {
	input := 2
	assert.Equal(t, 1, Fibonacci(input))
	input = 3
	assert.Equal(t, 2, Fibonacci(input))
	fmt.Println(Fibonacci(5))
	fmt.Println(Fibonacci(16))
	fmt.Println(Fibonacci(10))
}
func TestFibonacciRec(t *testing.T) {
	input := 2
	memo := make([]int, input+1)
	assert.Equal(t, 1, FibonacciRecurssive(input, memo))
	input = 3
	memo = make([]int, input+1)
	assert.Equal(t, 2, FibonacciRecurssive(input, memo))
	memo = make([]int, 5+1)
	fmt.Println(FibonacciRecurssive(5, memo))
	memo = make([]int, 16+1)
	fmt.Println(FibonacciRecurssive(16, memo))
	memo = make([]int, 100+1)
	fmt.Println(FibonacciRecurssive(100, memo))
	for i, v := range memo {
		fmt.Println(i, v)
	}
}

func TestClimbStaris(t *testing.T) {
	input := 2
	assert.Equal(t, 2, climbStairs(input))
	input = 3
	assert.Equal(t, 3, climbStairs(input))
	fmt.Println(climbStairs(5))
	fmt.Println(climbStairs(6))
}

func TestSubstringwithoutrepeatingchars(t *testing.T) {
	result := substringwithoutrepeatingchars("abcabc", 3)
	fmt.Println(result)
	assert.Equal(t, 3, len(result))
	result = substringwithoutrepeatingchars("abacab", 3)
	fmt.Println(result)
	assert.Equal(t, 2, len(result))
	result = substringwithoutrepeatingchars("awaglknagawunagwkwagl", 4)
	fmt.Println(result)
	assert.Equal(t, 12, len(result))
}

func TestSplit(t *testing.T) {
	input := "hello/world/test"
	result := strings.Split(input, "/")
	for _, v := range result {
		fmt.Println(v)
	}

}

func TestSlice(t *testing.T) {
	slice := make([]int, 10)
	fmt.Println(slice, len(slice), cap(slice))
	slice[5] = 100
	fmt.Println(slice[0:6])
	slice = append(slice, 200)
	slice = append(slice, 400)
	slice = append(slice, 400)
	fmt.Println(slice, len(slice), cap(slice))
}

func TestRobotInGrid(t *testing.T) {
	var row, col = 3, 3
	result := make([][]int, row)
	for i := 0; i < row; i++ {
		result[i] = make([]int, col)
	}
	assert.Equal(t, 6, RobotInGrid(row, col, result))
	row, col = 4, 4
	result = make([][]int, row)
	for i := 0; i < row; i++ {
		result[i] = make([]int, col)
	}
	fmt.Println(RobotInGrid(row, col, result))
}

func TestRobotInGridRecursiove(t *testing.T) {
	var row, col = 3, 3
	result := make([][]int, row)
	for i := 0; i < row; i++ {
		result[i] = make([]int, col)
	}

	assert.Equal(t, 6, RobotInGridRecursive(row-1, col-1, result))
	row, col = 4, 4
	result = make([][]int, row)
	for i := 0; i < row; i++ {
		result[i] = make([]int, col)
	}
	fmt.Println(RobotInGrid(row, col, result))

	row, col = 3, 7
	result = make([][]int, row)
	for i := 0; i < row; i++ {
		result[i] = make([]int, col)
	}
	fmt.Println(RobotInGridRecursive(row-1, col-1, result))
}

func TestSumofNumber(t *testing.T) {
	assert.Equal(t, 6, SumofNumber(12345))
	assert.Equal(t, 2, SumofNumber(749))
}

// TestfindTripplets ...
func TestFindTripplets(t *testing.T) {
	arr := []int{0, -1, 2, -3, 1}
	FindTripplets(arr, -2)
}

func TestFindTripplets2(t *testing.T) {
	arr := []int{5, 1, 3, 4, 7}
	FindTripplets2(arr, 12)
}

func TestPartition(t *testing.T) {
	// Input: set of integers
	S := []int{7, 3, 5, 12, 2, 1, 5, 3, 8, 4, 6, 4}
	k := 5
	Partition(S, len(S), k)
}

func TestSplitStringMaxSubstring(t *testing.T) {
	assert.Equal(t, 5, SplitStringMaxSubstring("ababccc"))
}

func TestRatInMaze(t *testing.T) {
	maze := [][]int{
		{1, 0, 0, 0},
		{1, 1, 0, 1},
		{0, 1, 0, 0},
		{1, 1, 1, 1}}
	RatInAMaze(maze, len(maze[0]), len(maze[0]))
}

func TestFindShortestPathMatrix(t *testing.T) {
	maze := [][]int{
		{1, 3, 1, 1},
		{1, 5, 1, 2},
		{4, 2, 1, 3},
		{6, 1, 2, 4},
	}
	assert.Equal(t, 13, FindShortestPathMatrix(maze))
}

func TestLongestPalindrome(t *testing.T) {
	//assert.Equal(t, "racecar", LongestPalindrome("racecar"))
	assert.Equal(t, "bab", LongestPalindrome("babcd"))
}
func TestLongestPalindrome2(t *testing.T) {
	//assert.Equal(t, 7, LongestPalindrome2("racecar"))
	assert.Equal(t, 3, LongestPalindrome2("bbacd"))
}

func TestLengthOfLongestSubstring(t *testing.T) {
	assert.Equal(t, 3, LengthOfLongestSubstring("abcabcbb"))
	assert.Equal(t, 1, LengthOfLongestSubstring("bbbbb"))
}

func TestLengthSubsetString(t *testing.T) {
	SubsetString("abc")
}

func TestLengthSubsets(t *testing.T) {
	input := []int{1, 2, 3}
	res := Subsets(input)
	for k, v := range res {
		fmt.Println(k, v)
	}
}

func TestNumIslands(t *testing.T) {
	grid := [][]byte{
		{'1', '1', '0', '0'},
		{'1', '1', '0', '1'},
		{'1', '1', '0', '1'},
		{'1', '1', '0', '0'},
	}
	fmt.Println(NumIslands(grid))
}

// Find All Numbers Disappeared in an Array
// Input:
// [4,3,2,7,8,2,3,1]
// Output:
// [5,6]

func TestFindDisappearedNumbers(t *testing.T) {
	input := []int{4, 3, 2, 7, 8, 2, 3, 1}
	result := FindDisappearedNumbers(input)
	fmt.Println(result)
}

func TestFindSquare(t *testing.T) {
	input := []int{-4, -1, 3, 5, 10}
	result := FindSquare(input)
	fmt.Println(result)
}

func TestMergeIntervals(t *testing.T) {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	result := MergeIntervals(intervals)
	fmt.Println(result)
}
