package functional_test

import (
	"testing"

	"github.com/JoshuaGanter/go-concurrency/functional"
	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	tests := []struct {
		input    []int
		filterFn func(int) bool
		expected []int
	}{
		{
			input:    []int{1, 2, 3, 4, 5},
			filterFn: func(elem int) bool { return elem < 4 },
			expected: []int{1, 2, 3},
		},
		{
			input:    []int{},
			filterFn: func(elem int) bool { return elem < 4 },
			expected: []int{},
		},
	}

	for _, test := range tests {
		var actual []int = functional.Filter(test.input, test.filterFn)

		assert.ElementsMatch(t, test.expected, actual)
	}
}
