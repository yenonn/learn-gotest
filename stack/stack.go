package stack

type Stack struct {
	data []string
}

func NewStack() *Stack {
	return &Stack{
		data: make([]string, 0),
	}
}

func (s *Stack) Size() int {
	return len(s.data)
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) <= 0
}

func (s *Stack) Push(input string) {
	s.data = append(s.data, input)
}

func (s *Stack) Pop() string {
	popItem := s.data[len(s.data)-1]
	// remove the popItem from the slice
	s.data = s.data[:len(s.data)-1]
	return popItem
}
