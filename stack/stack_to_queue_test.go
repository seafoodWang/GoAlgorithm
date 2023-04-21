package stack

import (
	"fmt"
	"testing"
)

func TestMyQueue(t *testing.T) {
	myQueue := ConstructorMyQueue()
	myQueue.Push(1)
	myQueue.Push(2)
	fmt.Println(myQueue.Pop())
	myQueue.Push(3)
	myQueue.Push(4)
	fmt.Println(myQueue.Pop())
	fmt.Println(myQueue.Pop())
	fmt.Println(myQueue.Peek())
	fmt.Println(myQueue.Pop())
}

type MyQueue struct {
	input  *Stack
	output *Stack
}

func ConstructorMyQueue() MyQueue {
	return MyQueue{InitStack(), InitStack()}
}

func (m *MyQueue) Push(x int) {
	m.input.Push(x)
}

func (m *MyQueue) Pop() int {
	if !m.output.IsEmpty() {
		d, _ := m.output.Pop()
		return d
	}
	InputToOutput(m.input, m.output)
	pop, _ := m.output.Pop()
	return pop
}

func (m *MyQueue) Peek() int {
	if !m.output.IsEmpty() {
		d, _ := m.output.Peek()
		return d
	}
	InputToOutput(m.input, m.output)
	peek, _ := m.output.Peek()
	return peek
}

func (m *MyQueue) Empty() bool {
	return m.input.IsEmpty() && m.output.IsEmpty()
}

func InputToOutput(input, output *Stack) {
	for !input.IsEmpty() {
		i, _ := input.Pop()
		output.Push(i)
	}
}
