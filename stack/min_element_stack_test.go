package stack_test

import (
	"fmt"
	stack "learn/stack"
	"testing"
)

// leetcode-155 https://leetcode.cn/problems/min-stack/
func TestGetMinElementStack(t *testing.T) {
	minElementStack := stack.InitMinElementStack()
	minElementStack.Push(7)
	minElementStack.Push(2)
	minElementStack.Push(5)
	minElementStack.Push(1)
	//1
	fmt.Println(minElementStack.GetMinElement())
	minElementStack.Pop()
	//2
	fmt.Println(minElementStack.GetMinElement())
	minElementStack.Push(10)
	//2
	fmt.Println(minElementStack.GetMinElement())
}
