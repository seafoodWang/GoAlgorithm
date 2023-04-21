package stack

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	stack := InitStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(5)
	fmt.Println(stack.Data)
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	stack.Push(10)
	fmt.Println(stack.Data)
	fmt.Println(stack.Pop())
	fmt.Println(stack.Peek())
	fmt.Println(stack.Data)
}
