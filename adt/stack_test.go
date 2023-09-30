package adt_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/yenonn/learn-gotest/adt"
)

type StackSuite struct {
	suite.Suite
}

func TestStackSuite(t *testing.T) {
	suite.Run(t, new(StackSuite))
}

func (s *StackSuite) TestEmpty() {
	stack := adt.NewStack()
	s.True(stack.IsEmpty())
}

func (s *StackSuite) TestNotEmpty() {
	stack := adt.NewStack()
	stack.Push("red")
	s.False(stack.IsEmpty())
}
