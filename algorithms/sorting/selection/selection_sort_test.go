package sorting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SelectionSortByMaxIndex(t *testing.T) {
	type Test struct {
		Given    []int64
		Expected []int64
	}

	tests := []Test{
		{
			[]int64{8, 5, 6, 4, 1, 2},
			[]int64{1, 2, 4, 5, 6, 8},
		},
		{
			[]int64{11, 2, 56, 8, 89},
			[]int64{2, 8, 11, 56, 89},
		},
	}

	for _, test := range tests {
		SelectionSortByMaxIndex(test.Given)
		assert.EqualValues(t, test.Expected, test.Given)
	}
}
