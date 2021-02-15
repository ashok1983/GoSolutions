package main

// Coding Skill Exercise #1
//  FibonacciRecurssive
func FibonacciRec(n int, memo []int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else if memo[n] > 0 {
		return memo[n]
	}

	memo[n] = FibonacciRec(n-1, memo) + FibonacciRec(n-2, memo)

	return memo[n]
}

// Coding Skill Exercise #2
// Compute a % b without using % operator.

func computeMod(a int, b int) int {

	rem := a / b
	prod := rem * b
	result := a - prod
	return result
}

// Coding Skill Exercise #3
// Compute a to the power b.

func computePowerofNum(a int, b int) int {
	result := 1
	for ; b != 0; b-- {
		result *= a
	}
	return result
}
