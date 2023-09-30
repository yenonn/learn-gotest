package adt

type Stack struct {
	data []string
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) <= 0
}

func (s *Stack) Push(input string) {
	s.data = append(s.data, input)
}

func (s *Stack) Size() int {
	return len(s.data)
}
