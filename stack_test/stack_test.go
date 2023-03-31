package stack

import (
	"testing"
)

func TestStack(t *testing.T) {
	stack := InitStack(4)
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(5)
	stack.IsFull()

}
