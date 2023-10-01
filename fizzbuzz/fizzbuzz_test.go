package fizzbuzz_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/yenonn/learn-gotest/fizzbuzz"
)

type FizzBuzzSuite struct {
	suite.Suite
}

func TestFizzBuzz(t *testing.T) {
	suite.Run(t, new(FizzBuzzSuite))
}

func (f *FizzBuzzSuite) TestOneUnchanged() {
	r := fizzbuzz.Run([]int{1})
	f.Equal([]string{"1"}, r)
}
