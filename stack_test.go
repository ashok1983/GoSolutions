package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	st := &mystack{}
	st.push(10)
	st.push(20)
	st.push(30)
	st.push(40)

	st.printstack()
	st.pop()
	st.pop()
	st.pop()
	st.pop()
	st.pop()
	st.printstack()
}

func TestValidParenthesic(t *testing.T) {
	par := "()"
	assert.Equal(t, true, IsValidParenteses(par))
	par = "()[]"
	assert.Equal(t, true, IsValidParenteses(par))
	par = "(((())))"
	assert.Equal(t, true, IsValidParenteses(par))
	par = "()[]{}"
	assert.Equal(t, true, IsValidParenteses(par))
	par = "([])"
	assert.Equal(t, true, IsValidParenteses(par))
	par = "(][])"
	assert.Equal(t, false, IsValidParenteses(par))
	par = "(((()"
	assert.Equal(t, false, IsValidParenteses(par))
	par = "())("
	assert.Equal(t, false, IsValidParenteses(par))
}
