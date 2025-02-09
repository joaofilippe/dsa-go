package sorting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickSorting(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{input: []int{3, 6, 8, 10, 1, 2, 5}, expected: []int{1, 2, 3, 5, 6, 8, 10}},
		{input: []int{1, 2, 3, 4, 5}, expected: []int{1, 2, 3, 4, 5}},
		{input: []int{5, 4, 3, 2, 1}, expected: []int{1, 2, 3, 4, 5}},
		{input: []int{}, expected: []int{}},
		{input: []int{1}, expected: []int{1}},
	}

	for _, test := range tests {
		result := QuickSorting(test.input)
		assert.Equal(t, test.expected, result)
	}
}
