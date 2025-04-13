package functional_test

import (
	"testing"

	"github.com/JoshuaGanter/go-concurrency/functional"
	"github.com/stretchr/testify/assert"
)

type TestCase[TValue comparable] struct {
	input    []TValue
	filterFn func(TValue, int, []TValue) bool
	expected []TValue
}

func TestFilter(t *testing.T) {
	intValueTests := []TestCase[int]{
		{
			input:    []int{1, 2, 3, 4, 5},
			filterFn: func(element int, index int, slice []int) bool { return element < 4 },
			expected: []int{1, 2, 3},
		},
		{
			input:    []int{},
			filterFn: func(element int, index int, slice []int) bool { return element < 4 },
			expected: []int{},
		},
		{
			input:    []int{42, 84, 1, 5, 77, 11},
			filterFn: func(element int, index int, slice []int) bool { return index%2 == 0 },
			expected: []int{42, 1, 77},
		},
	}

	for _, test := range intValueTests {
		actual := functional.Filter(test.input, test.filterFn)

		assert.ElementsMatch(t, test.expected, actual)
	}

	stringValueTests := []TestCase[string]{
		{
			input:    []string{"hello", "World", "from", "Inside", "GO"},
			filterFn: func(element string, index int, slice []string) bool { return len(element) == 5 },
			expected: []string{"hello", "World"},
		},
	}

	for _, test := range stringValueTests {
		actual := functional.Filter(test.input, test.filterFn)

		assert.ElementsMatch(t, test.expected, actual)
	}
}
