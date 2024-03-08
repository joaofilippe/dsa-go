package sorting

import (
	"testing"

	"github.com/joaofilippe/dsa-go/enums"
	"github.com/stretchr/testify/suite"
)

type SuiteBubbleSorting struct {
	suite.Suite
	slices   [][]int64
	expected [][]int64
}

func (suite *SuiteBubbleSorting) SetupSuite() {
	suite.slices = [][]int64{
		{5, 4, 2, 3, 7},
		{5, 3, 1, 2, 4},
		{5, 1, 7, 3, 2},
		{5, 3, 1, 8, 0},
		{5, 2, 12, 9, 7},
	}
}

func (suite *SuiteBubbleSorting) TearDownTest() {
	suite.expected = nil
}

func (suite *SuiteBubbleSorting) Test_BubbleSortingGreaterThan() {
	suite.expected = [][]int64{
		{2, 3, 4, 5, 7},
		{1, 2, 3, 4, 5},
		{1, 2, 3, 5, 7},
		{0, 1, 3, 5, 8},
		{2, 5, 7, 9, 12},
	}

	for i, slice := range suite.slices {
		InsertionSort(slice, enums.GREATER_THAN)
		suite.Equal(suite.expected[i], slice)
	}
}

func (suite *SuiteBubbleSorting) Test_BubbleSortingLessThan() {
	suite.expected = [][]int64{
		{7, 5, 4, 3, 2},
		{5, 4, 3, 2, 1},
		{7, 5, 3, 2, 1},
		{8, 5, 3, 1, 0},
		{12, 9, 7, 5, 2},
	}

	for i, slice := range suite.slices {
		InsertionSort(slice, enums.LESS_THAN)
		suite.Equal(suite.expected[i], slice)
	}
}

func Test_BubleSorting(t *testing.T) {
	suite.Run(t, new(SuiteBubbleSorting))
}