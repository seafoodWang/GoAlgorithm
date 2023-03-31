package stack

import "errors"

type MinElementStack struct {
	DataStack *Stack
	MinStack  *Stack
}

func InitMinElementStack(size int) *MinElementStack {
	dataStack := InitStack(size)
	minStack := InitStack(size)
	return &MinElementStack{DataStack: dataStack, MinStack: minStack}
}

func (s *MinElementStack) IsEmpty() bool {
	return s.DataStack.IsEmpty()
}

func (s *MinElementStack) IsFull() bool {
	return s.DataStack.IsFull()
}

func (s *MinElementStack) Push(x int) (int, error) {
	if !s.IsFull() {
		s.DataStack.Push(x)
		if s.MinStack.IsEmpty() {
			s.MinStack.Push(x)
		} else {
			p, _ := s.MinStack.Peek()
			if p > x {
				p = x
			}
			s.MinStack.Push(p)
		}
	}

	return x, errors.New("is full")
}

func (s *MinElementStack) Pop() (int, error) {
	if s.IsEmpty() {
		return -1, errors.New("stack is Empty can not pop")
	}
	x, _ := s.DataStack.Pop()
	s.MinStack.Pop()
	return x, nil
}

func (s *MinElementStack) GetMinElement() int {
	p, _ := s.MinStack.Peek()
	return p
}
