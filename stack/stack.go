package stack

import (
	"errors"
)

type Stack struct {
	Data []int
}

// 动态扩容
func InitStack() *Stack {
	data := make([]int, 0)
	return &Stack{
		Data: data,
	}
}

func (s *Stack) IsEmpty() bool {
	return len(s.Data) == 0
}

func (s *Stack) IsFull() bool {
	return len(s.Data) == cap(s.Data)
}

func (s *Stack) Push(x int) int {
	s.Data = append(s.Data, x)
	return x
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is Empty can not pop")
	}
	x := s.Data[len(s.Data)-1]
	s.Data = s.Data[:len(s.Data)-1]
	return x, nil
}

func (s *Stack) Peek() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("is empty can not peek")
	}
	return s.Data[len(s.Data)-1], nil
}
