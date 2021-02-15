package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacciRecAssign(t *testing.T) {
	input := 2
	memo := make([]int, input+1)
	assert.Equal(t, 1, FibonacciRec(input, memo))
	input = 3
	memo = make([]int, input+1)
	assert.Equal(t, 2, FibonacciRec(input, memo))
	memo = make([]int, 5+1)
	fmt.Println(FibonacciRec(5, memo))
	memo = make([]int, 16+1)
	fmt.Println(FibonacciRec(16, memo))
	memo = make([]int, 100+1)
	fmt.Println(FibonacciRec(100, memo))
	for i, v := range memo {
		fmt.Println(i, v)
	}
}

func TestMod(t *testing.T) {
	assert.Equal(t, 2, computeMod(27, 5))
	assert.Equal(t, 0, computeMod(8, 2))
}

func TestPower(t *testing.T) {
	assert.Equal(t, 256, computePowerofNum(2, 8))
	assert.Equal(t, 1, computePowerofNum(2, 0))
}
