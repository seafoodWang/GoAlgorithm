package stack

import "errors"

type Stack struct {
	Size  int
	Data  []int
	index int
}

func InitStack(size int) *Stack {
	data := make([]int, size)
	return &Stack{
		Size:  size,
		Data:  data,
		index: 0,
	}
}

func (s *Stack) IsEmpty() bool {
	return s.index == 0
}

func (s *Stack) IsFull() bool {
	return s.Size == s.index
}

func (s *Stack) Push(x int) (int, error) {
	if s.IsFull() {
		return -1, errors.New("stack is Full can not push")
	}
	s.Data[s.index] = x
	s.index += 1
	return x, nil
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return -1, errors.New("stack is Empty can not pop")
	}
	s.index -= 1
	x := s.Data[s.index]
	s.Data[s.index] = 0
	return x, nil
}

func (s *Stack) Peek() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("is empty can not peek")
	}
	i := s.index - 1
	return s.Data[i], nil
}
