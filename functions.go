package main

import (
	"fmt"
	"sort"
	"strings"
)

// RemoveDuplicate ....
func RemoveDuplicate(input []int, len int) {
	if len == 0 {
		fmt.Println("lenght is zero")
		return
	}
	result := make(map[int]bool)
	for i := 0; i < len; i++ {
		if result[input[i]] != true {
			result[input[i]] = true
		}
	}
	fmt.Println(result)
	var output []int
	for k := range result {
		output = append(output, k)
	}
	sort.Ints(output)
	fmt.Println(output)
	return
}

// RemoveDuplicateSameArray ....
func RemoveDuplicateSameArray(input []int, len int) {
	if len == 0 {
		fmt.Println("invalid input ")
		return
	}

	var index int
	for i := 0; i < len; i++ {
		if i < len-1 && input[i] == input[i+1] {
			continue
		}
		input[index] = input[i]
		index++
	}

	fmt.Println(input)
}

// IsRotation ...
func IsRotation(s1 string, s2 string) bool {

	if s1 == "" || s2 == "" {
		fmt.Println("Invalid input")
		return false
	}

	s3 := s1 + s1
	return strings.Contains(s3, s2)
}

func isUniqueChars(s1 string) bool {
	if len(s1) == 0 {
		return false
	}

	result := make(map[string]bool)
	for _, val := range s1 {
		if result[string(val)] == true {
			fmt.Println("not unique string : ", s1)
			return false
		}
		result[string(val)] = true
	}
	fmt.Println("unique string :", s1)
	return true
}

//Ispermutation ...
func Ispermutation(s1 string, s2 string) bool {
	if len(s1) != len(s2) && s1 == "" || s2 == "" {
		fmt.Println("not permutation")
		return false
	}
	result := make(map[string]int)

	for _, val := range s1 {
		result[string(val)]++
	}

	for _, val := range s2 {
		result[string(val)]--
		if result[string(val)] < 0 {
			fmt.Println("not a permutation ")
			return false
		}
	}
	fmt.Println("permutation")
	return true
}

//Ispermutation2 ...
func Ispermutation2(s1 string, s2 string) bool {
	if len(s1) != len(s2) && s1 == "" || s2 == "" {
		fmt.Println("not permutation")
		return false
	}

	result := make(map[byte]int)

	for _, val := range s1 {
		result[byte(val)]++
	}

	for _, val := range s2 {
		result[byte(val)]--
		if result[byte(val)] < 0 {
			fmt.Println("not a permutation ")
			return false
		}
	}
	fmt.Println("permutation")
	return true
}

//A palindrome is a string that is the same forwards and backwards. Therefore, to decide if a string is a permutation
//of a palindrome, we need to know if it can be written such that it's the same forwards and backwards

// IspermutationOfPalindrome ...
func IspermutationOfPalindrome(s1 string) bool {

	if s1 == "" {
		fmt.Println("invalid input")
		return false
	}

	// make map of occurance of charactes.
	charMap := make(map[rune]int)

	for _, c := range s1 {
		charMap[c]++
	}
	fmt.Println(charMap)
	var oddfound bool
	// check the occurance of character is all even and only one ever
	for _, v := range charMap {
		if v%2 == 1 {
			if oddfound == true {
				return false
			}
			oddfound = true
		}

	}
	return true
}

func generatePermutation(s1 []rune, left int, right int) {
	if left == right {
		fmt.Println(string(s1))
	} else {
		for i := left; i <= right; i++ {
			s1[left], s1[i] = s1[i], s1[left]
			generatePermutation(s1, left+1, right)
			s1[left], s1[i] = s1[i], s1[left]
		}
	}
}

func generatePermutationByte(s1 []byte, left int, right int) {
	if left == right {
		fmt.Println(string(s1))
	} else {
		for i := left; i <= right; i++ {
			s1[left], s1[i] = s1[i], s1[left]
			fmt.Println("above recursion", i, left, string(s1))
			generatePermutationByte(s1, left+1, right)
			s1[left], s1[i] = s1[i], s1[left]
			fmt.Println("below recursion", i, left, string(s1))
		}
	}
}

func swap(s1 *string, i int, j int) {
	input := []rune(*s1)
	input[i], input[j] = input[j], input[i]
}

// subset of string...
func SubsetString(s string) [][]string {
	result := make([][]string, 1)

	for i := 0; i < len(s); i++ {
		for _, v := range result {
			result = append(result, append(v, string(s[i])))
		}
	}

	for k, v := range result {
		fmt.Println(k, v)
	}
	return result
}

// Subsets ...
func Subsets(nums []int) [][]int {
	res := make([][]int, 1)
	for _, v := range nums {
		for _, r := range res {
			res = append(res, append(r, v))
		}
	}
	return res
}

func balancedString(str string) int {
	res, count := 0, 0
	for _, c := range str {
		if c == 'L' {
			count++
		} else {
			count--
		}
		if count == 0 {
			res++
		}
	}
	return res
}

// dynamic programming
func climbStairs(n int) int {
	var memo = []int{1, 1}
	for i := 2; i <= n; i++ {
		memo = append(memo, memo[i-1]+memo[i-2])
	}
	return memo[n]
}

// dynamic programming ...
func Fibonacci(n int) int {
	var memo = []int{0, 1}
	for i := 2; i <= n; i++ {
		memo = append(memo, memo[i-1]+memo[i-2])
	}
	return memo[n]
}

//  FibonacciRecurssive
func FibonacciRecurssive(n int, memo []int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else if memo[n] > 0 {
		return memo[n]
	}

	memo[n] = FibonacciRecurssive(n-1, memo) + FibonacciRecurssive(n-2, memo)

	return memo[n]
}

// input : string
// k  window
func substringwithoutrepeatingchars(input string, k int) (result []string) {
	if input == "" {
		fmt.Println("invalid input")
		return nil
	}
	output := make(map[string]bool)
	var valstr string
	for i := 0; i < len(input); i++ {
		if (i + k) <= len(input) {
			valstr = input[i:(i + k)]
			if validstring(valstr) && output[valstr] != true {
				output[valstr] = true
			}
		}
	}

	for k := range output {
		result = append(result, k)
	}
	return result
}
func validstring(input string) bool {
	result := make(map[rune]bool)

	for _, c := range input {
		if result[c] == true {
			return false
		}
		result[c] = true
	}
	return true
}

// find the total number of paths
//Problem: A robot is located at the top-left corner of a m x n grid.
//The robot can only move either down or right at any point in time. The robot is trying to reach the bottom-right corner of the grid .
// How many possible unique paths are there?
// apporach1 : 1 dyanmic programing approach (diagonal addition)
// 1 1 1
// 1 2 3
// 1 3 6

// RobotInGrid ...
func RobotInGrid(r int, c int, result [][]int) int {
	for i := 0; i < r; i++ {
		result[i][0] = 1
	}
	for j := 0; j < c; j++ {
		result[0][j] = 1
	}
	for i := 1; i < r; i++ {
		for j := 1; j < c; j++ {
			result[i][j] = result[i-1][j] + result[i][j-1]
		}
	}
	fmt.Println(result)
	return result[r-1][c-1]
}

//backtrack
func RobotInGridRecursive(row int, col int, memo [][]int) int {

	if row == 0 || col == 0 {
		return 1
	}

	if memo[row][col] != 0 {
		return memo[row][col]
	} else {
		memo[row][col] = RobotInGridRecursive(row-1, col, memo) + RobotInGridRecursive(row, col-1, memo)
		return memo[row][col]
	}

}

//Program to find the sum of the digits of a number untill the sum is reduced to a single digit
// EX: 12345 result 6
// 749 : 2

// SumofNumber ...
func SumofNumber(num int) int {
	fmt.Println(num)
	var sum int
	var rem int

	for num/10 != 0 {
		sum = 0
		for ; num != 0; num = num / 10 {
			rem = num % 10
			sum += rem
		}
		num = sum
	}
	return sum
}

//Input: arr[] = {0, -1, 2, -3, 1}
// sum = -2
// Output:  0 -3  1
//         -1  2 -3
// If we calculate the sum of the output,
// 0 + (-3) + 1 = -2
// (-1) + 2 + (-3) = -2

// Input: arr[] = {1, -2, 1, 0, 5}
//        sum = 0
// Output: 1 -2  1
// If we calculate the sum of the output,
// 1 + (-2) + 1 = 0

func FindTripplets(input []int, sum int) {

	size := len(input)

	for i := 0; i < size-1; i++ {
		result := make(map[int]bool)
		for j := i + 1; j < size; j++ {
			x := sum - (input[i] + input[j])
			if result[x] == true {
				fmt.Println(input[i], input[j], x)
			}
			result[input[j]] = true
		}
	}
}

// FindTripplets2 ...
func FindTripplets2(input []int, sum int) {
	size := len(input)
	sort.Ints(input)

	for i := 0; i < size-2; i++ {
		begin := i + 1
		end := size - 1
		for begin < end {
			if input[i]+input[begin]+input[end] < sum {
				//fmt.Println(input[i], input[begin], input[end])
				// This is important. For current i and j,
				// there are total k-j third elements.
				for x := begin + 1; x <= end; x++ {
					fmt.Println(input[i], input[begin], input[x])
				}

				begin++
			} else {
				end--
			}
		}
	}
}

// K partition program
// Function for solving k-partition problem. It prints the subsets if
// set S[0..n-1] can be divided into k subsets with equal sum
func Partition(input []int, size int, k int) {
	if size < k {
		fmt.Println("invalid array")
		return
	}

	sum := 0
	for i := 0; i < size; i++ {
		sum += input[i]
	}

	A := make([]int, size)
	SumLeft := make([]int, k)

	for i := 0; i < k; i++ {
		SumLeft[i] = sum / k
	}

	res := helper(input, size-1, SumLeft, A, k)

	if res {
		// print all k-partitions
		for i := 0; i < k; i++ {
			for j := 0; j < size; j++ {
				if A[j] == i+1 {
					fmt.Printf("%d ", input[j])
				}
			}
			fmt.Println()
		}
	}
}

func checksum(SumLeft []int, k int) bool {
	r := true
	for i := 0; i < k; i++ {
		if SumLeft[i] != 0 {
			r = false
		}
	}
	return r
}
func helper(input []int, size int, SumLeft []int, A []int, k int) bool {

	if checksum(SumLeft, k) {
		return true
	}

	if size < 0 {
		return false
	}

	var res bool
	for i := 0; i < k; i++ {
		if !res && SumLeft[i]-input[size] >= 0 {
			A[size] = i + 1
			SumLeft[i] = SumLeft[i] - input[size]
			res = helper(input, size-1, SumLeft, A, k)
			SumLeft[i] = SumLeft[i] + input[size]
		}

	}
	return res
}

// Split a String Into the Max Number of Unique Substrings ...
// Input: s = "ababccc"
// Output: 5
// Explanation: One way to split maximally is ['a', 'b', 'ab', 'c', 'cc'].
// Splitting like ['a', 'b', 'a', 'b', 'c', 'cc'] is not valid as you have 'a' and 'b' multiple times.
func SplitStringMaxSubstring(input string) int {
	if len(input) == 0 {
		fmt.Println("Invalid lenght")
		return 0
	}
	result := make(map[string]bool)
	res := dfs(input, result, 0)
	fmt.Println(result)
	return res
}

func dfs(input string, result map[string]bool, i int) int {
	var res int
	if len(input) == i {
		return 0
	}

	for j := i; j < len(input); j++ {
		str := input[i : j-i+1]
		fmt.Println(str)
		if result[str] {
			continue
		}
		result[str] = true
		res = max(res, 1+dfs(input, result, j+1))
		delete(result, str)
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

//Rat in a Maze | Backtracking-2

var result = [][]int{{0, 0, 0, 0},
	{0, 0, 0, 0},
	{0, 0, 0, 0},
	{0, 0, 0, 0}}

func RatInAMaze(maze [][]int, row int, col int) bool {
	if row == 0 || col == 0 {
		fmt.Println("Invalid metrics")
		return false
	}

	MazeHelper(maze, row, col, 0, 0)
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			fmt.Printf("%d ", result[i][j])
		}
		fmt.Println()
	}
	return true
}

// MazeHelper ...
func MazeHelper(maze [][]int, r int, c int, x int, y int) bool {

	// condition for success goal
	if x == r-1 && y == c-1 && maze[x][y] == 1 {
		result[x][y] = 1
		return true
	}

	if isSafe(maze, r, c, x, y) == true {
		result[x][y] = 1
		if MazeHelper(maze, r, c, x+1, y) == true {
			return true
		}
		if MazeHelper(maze, r, c, x, y+1) == true {
			return true
		}
		// backtrack
		result[x][y] = 0
		return false
	}

	return false
}

func isSafe(maze [][]int, r int, c int, x int, y int) bool {
	if x >= 0 && x < r && y >= 0 && y < c && maze[x][y] == 1 {
		return true
	}
	return false
}

// Find the shortest path to reach destination

func FindShortestPathMatrix(input [][]int) int {

	m := len(input)
	n := len(input[0])

	if m == 0 && n == 0 {
		return 0
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if i == 0 {
				input[i][j] += input[i][j-1]
			} else if j == 0 {
				input[i][j] += input[i-1][j]

			} else {
				input[i][j] += min(input[i][j-1], input[i-1][j])
			}
		}
	}
	return input[m-1][n-1]
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

// LongestPalindrome ....
//Input: s = "babad"
//Output: "bab"
//Note: "aba" is also a valid answer.

func LongestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	var start int
	var end int

	for i := 0; i < len(s); i++ {
		len1 := expandFromMiddle(s, i, i)
		len2 := expandFromMiddle(s, i, i+1)
		size := max(len1, len2)
		if size > end-start {
			start = i - (size-1)/2
			end = i + size/2
		}
	}
	return s[start : end+1]
}

func expandFromMiddle(s string, left int, right int) int {
	if len(s) == 0 || left > right {
		return 0
	}

	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	fmt.Println((right - left - 1))
	return (right - left - 1)
}

func LongestPalindrome2(s string) int {
	var count [25]int
	for i := range s {
		count[s[i]-'a']++
	}

	ans := 0
	fmt.Println(count)
	for _, v := range count {
		ans += v / 2 * 2
		fmt.Println(ans)
		if ans%2 == 0 && v%2 == 1 {
			ans++
		}
	}

	return ans

}

//  abcabcbb   : 3 abc
// bbbbb : 1 b
func LengthOfLongestSubstring(s string) int {
	size := len(s)
	result := make(map[byte]bool)
	var ans int
	var i int
	var j int

	for i < size && j < size {
		if result[s[j]] == false {
			result[s[j]] = true
			j++
			ans = max(ans, j-i)
		} else {
			delete(result, s[i])
			i++
		}

	}
	return ans
}

func NumIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	var island int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				island++
				dfsIsland(grid, i, j)
			}
		}
	}

	return island
}

func dfsIsland(grid [][]byte, x int, y int) {
	if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) || grid[x][y] == '0' {
		return
	}
	fmt.Println(grid)
	grid[x][y]++
	fmt.Println(grid)
	dfsIsland(grid, x+1, y)
	dfsIsland(grid, x, y+1)
	dfsIsland(grid, x-1, y)
	dfsIsland(grid, x, y-1)
}

// func FindDisappearedNumbers(nums []int) []int {

// 	nlen := len(nums)
// 	for i := 0; i < nlen; i++ {
// 		nums[(nums[i]-1)%nlen] += nlen
// 	}

// 	out := []int{}
// 	for i := 0; i < nlen; i++ {
// 		if nums[i] <= nlen {
// 			out = append(out, i+1)
// 		}
// 	}
// 	return out
// }

func FindDisappearedNumbers(nums []int) []int {
	result := []int{}
	if len(nums) == 0 {
		return result
	}
	output := make(map[int]bool, len(nums))
	for _, v := range nums {
		output[v] = true
	}
	for i := 1; i <= len(nums); i++ {
		if output[i] == false {
			result = append(result, i)
		}
	}
	return result
}

// FindSquare
func FindSquare(input []int) []int {
	if len(input) == 0 {
		return input
	}

	for i, v := range input {
		input[i] = v * v
	}

	low := 0
	high := len(input) - 1
	for high > low {
		if input[low] > input[high] {
			input[high], input[low] = input[low], input[high]

		}
		high--
	}
	fmt.Println(input)
	return input
}

/*Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]
Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into [1,6].*/

func MergeIntervals(intervals [][]int) [][]int {
	var result [][]int
	if len(intervals) > 0 {
		//sort
		sort.Slice(intervals, func(i, j int) bool {
			return intervals[i][0] < intervals[j][0]
		})
		currentVal := intervals[0]
		for i := 1; i < len(intervals); i++ {
			if currentVal[1] >= intervals[i][0] {
				if currentVal[1] < intervals[i][1] {
					currentVal[1] = intervals[i][1]
				}
			} else {
				result = append(result, currentVal)
				currentVal = intervals[i]
			}
		}
		result = append(result, currentVal)
	}
	return result
}
