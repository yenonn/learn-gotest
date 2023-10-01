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
	result := fizzbuzz.Run([]int{1})
	f.Equal([]string{"1"}, result)
}

func (f *FizzBuzzSuite) TestTwoUnchanged() {
	result := fizzbuzz.Run([]int{1, 2})
	f.Equal([]string{"1", "2"}, result)
}

func (f *FizzBuzzSuite) TestThreeChanged() {
	result := fizzbuzz.Run([]int{1, 2, 3})
	f.Equal([]string{"1", "2", "fizz"}, result)
}
