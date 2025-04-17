package functional_test

import (
	"testing"

	"github.com/JoshuaGanter/go-functional/functional"
	"github.com/stretchr/testify/assert"
)

type reduceTestCase[TValue comparable, TResult comparable] struct {
	input    []TValue
	reduceFn func(TResult, TValue, int, []TValue) TResult
	initial  TResult
	expected TResult
}

func TestReduce(t *testing.T) {
	intToIntTests := []reduceTestCase[int, int]{
		{
			input:    []int{},
			reduceFn: func(previous int, current int, index int, slice []int) int { return previous + current },
			initial:  0,
			expected: 0,
		},
		{
			input:    []int{2, 4, 6, 8},
			reduceFn: func(previous int, current int, index int, slice []int) int { return previous + current },
			initial:  0,
			expected: 20,
		},
	}

	for _, test := range intToIntTests {
		actual := functional.Reduce(test.input, test.reduceFn, test.initial)

		assert.Equal(t, test.expected, actual)
	}

	stringToStringTests := []reduceTestCase[string, string]{
		{
			input:    []string{"hello", "World", "from", "Inside", "GO"},
			reduceFn: func(previous string, current string, index int, slice []string) string { return previous + current },
			initial:  "",
			expected: "helloWorldfromInsideGO",
		},
	}

	for _, test := range stringToStringTests {
		actual := functional.Reduce(test.input, test.reduceFn, test.initial)

		assert.Equal(t, test.expected, actual)
	}

	stringToIntTests := []reduceTestCase[string, int]{
		{
			input:    []string{"hello", "World", "from", "Inside", "GO"},
			reduceFn: func(previous int, current string, index int, slice []string) int { return previous + len(current) },
			initial:  0,
			expected: 22,
		},
	}

	for _, test := range stringToIntTests {
		actual := functional.Reduce(test.input, test.reduceFn, test.initial)

		assert.Equal(t, test.expected, actual)
	}
}
