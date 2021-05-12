package stack

import "fmt"

type Stack struct {
	data []interface{}
	top  int // 栈顶指针
}

func NewStack() *Stack {
	return &Stack{
		data: make([]interface{}, 0, 32),
		top:  -1,
	}
}

func (s *Stack) IsEmpty() bool {
	if s.top < 0 {
		return true
	}
	return false
}

func (s *Stack) Push(v interface{}) {
	if s.top < 0 {
		s.top = 0
	} else {
		s.top++
	}

	if s.top > len(s.data)-1 {
		s.data = append(s.data, v)
	} else {
		s.data[s.top] = v
	}
}

func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	v := s.data[s.top]
	s.top--
	return v
}

func (s *Stack) Top() interface{} {
	if s.IsEmpty() {
		return nil
	}
	return s.data[s.top]
}

func (s *Stack) Flush() {
	s.top = -1
}

func (s *Stack) Print() {
	if s.IsEmpty() {
		fmt.Println("empty stack")
	} else {
		for i := s.top; i >= 0; i-- {
			fmt.Println(s.data[i])
		}
	}
}
