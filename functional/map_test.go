package functional_test

import (
	"strings"
	"testing"

	"github.com/JoshuaGanter/go-functional/functional"
	"github.com/stretchr/testify/assert"
)

type mapTestCase[TValue comparable, TResult comparable] struct {
	input    []TValue
	mapFn    func(TValue, int, []TValue) TResult
	expected []TResult
}

func TestMap(t *testing.T) {
	intToIntTests := []mapTestCase[int, int]{
		{
			input:    []int{},
			mapFn:    func(element int, index int, slice []int) int { return element * element },
			expected: []int{},
		},
		{
			input:    []int{2, 4, 6, 8},
			mapFn:    func(element int, index int, slice []int) int { return element * element },
			expected: []int{4, 16, 36, 64},
		},
	}

	for _, test := range intToIntTests {
		actual := functional.Map(test.input, test.mapFn)

		assert.ElementsMatch(t, test.expected, actual)
	}

	stringToStringTests := []mapTestCase[string, string]{
		{
			input:    []string{"hello", "World", "from", "Inside", "GO"},
			mapFn:    func(element string, index int, slice []string) string { return strings.ToUpper(element) },
			expected: []string{"HELLO", "WORLD", "FROM", "INSIDE", "GO"},
		},
	}

	for _, test := range stringToStringTests {
		actual := functional.Map(test.input, test.mapFn)

		assert.ElementsMatch(t, test.expected, actual)
	}

	stringToIntTests := []mapTestCase[string, int]{
		{
			input:    []string{"hello", "World", "from", "Inside", "GO"},
			mapFn:    func(element string, index int, slice []string) int { return len(element) },
			expected: []int{5, 5, 4, 6, 2},
		},
	}

	for _, test := range stringToIntTests {
		actual := functional.Map(test.input, test.mapFn)

		assert.ElementsMatch(t, test.expected, actual)
	}
}
