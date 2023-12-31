package stack_test

import (
	"log"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/yenonn/learn-gotest/stack"
)

type StackSuite struct {
	suite.Suite
	stack *stack.Stack
}

func TestStackSuite(t *testing.T) {
	log.Println("Learning testify suite")
	suite.Run(t, new(StackSuite))
}

// Setup the Suite
// Put all of the setup lines here.
func (s *StackSuite) SetupTest() {
	s.stack = stack.NewStack()
}

func (s *StackSuite) TestEmpty() {
	s.True(s.stack.IsEmpty())
}

func (s *StackSuite) TestNotEmpty() {
	s.stack.Push("red")
	s.False(s.stack.IsEmpty())
}

func (s *StackSuite) TestStackEmptySizeZero() {
	s.Zero(s.stack.Size())
}

func (s *StackSuite) TestStackNotEmptySize() {
	s.stack.Push("red")
	s.stack.Push("blue")
	s.NotZero(s.stack.Size())
	s.Equal(2, s.stack.Size())
}

func (s *StackSuite) TestStackPop() {
	s.stack.Push("red")
	s.Equal("red", s.stack.Pop())
	s.True(s.stack.IsEmpty())
	s.Zero(s.stack.Size())
}

func (s *StackSuite) TestStackPopMultipleItems() {
	// FiLo, first in, last out
	s.stack.Push("red")
	s.stack.Push("blue")
	s.Equal("blue", s.stack.Pop())
	s.Equal("red", s.stack.Pop())
	s.Zero(s.stack.Size())
	s.True(s.stack.IsEmpty())
}
