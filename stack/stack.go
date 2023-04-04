package stack

import (
	"errors"
	"fmt"
)

type Stack struct {
	Size  int
	Data  []int
	index int
}

// 动态扩容
func InitStack() *Stack {
	data := make([]int, 0)
	return &Stack{
		Data:  data,
		index: 0,
	}
}

func (s *Stack) IsEmpty() bool {
	return len(s.Data) == 0
}

func (s *Stack) IsFull() bool {
	return len(s.Data) == cap(s.Data)
}

func (s *Stack) Push(x int) (int, error) {
	fmt.Println(s.Data)
	if s.index < cap(s.Data) {
		s.Data[s.index] = x
	} else {
		s.Data = append(s.Data, x)
	}
	s.index++
	return x, nil
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is Empty can not pop")
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
