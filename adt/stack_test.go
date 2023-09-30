package adt_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/yenonn/learn-gotest/adt"
)

type StackSuite struct {
	suite.Suite
	stack *adt.Stack
}

func TestStackSuite(t *testing.T) {
	suite.Run(t, new(StackSuite))
}

// Setup the Suite
// Put all of the setup lines here.
func (s *StackSuite) SetupTest() {
	s.stack = adt.NewStack()
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
	s.Zero(s.stack.Size())
}
