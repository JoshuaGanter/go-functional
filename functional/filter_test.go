package functional_test

import (
	"testing"

	"github.com/JoshuaGanter/go-concurrency/functional"
	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	tests := []struct {
		input    []int
		filterFn func(int, int, []int) bool
		expected []int
	}{
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

	for _, test := range tests {
		var actual []int = functional.Filter(test.input, test.filterFn)

		assert.ElementsMatch(t, test.expected, actual)
	}
}
