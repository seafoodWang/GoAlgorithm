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
}
