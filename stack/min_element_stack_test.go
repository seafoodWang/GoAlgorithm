package stack_test

import (
	"fmt"
	stack "learn/stack_test"
	"testing"
)

func TestGetMinElementStack(t *testing.T) {
	minElementStack := stack.InitMinElementStack(5)
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
