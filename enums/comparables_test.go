package enums

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type SuiteComparable struct {
	suite.Suite
	numbers    [][]int64
	comparable []Comparable
	expected   []bool
}

func (s *SuiteComparable) Test_ComparableFunc() {
	fmt.Println("Running test, TestComparableFunc")

	for i, v := range s.numbers {
		result := s.comparable[i].Compare(v[0], v[1])
		s.Equal(s.expected[i], result)
	}
}

func (s *SuiteComparable) SetupSuite() {
	fmt.Println("Starting suite")
}

func (s *SuiteComparable) TearDownSuite() {
	fmt.Println("Ending suite")
}

func (s *SuiteComparable) SetupTest() {
	fmt.Println("Starting test")

	s.numbers = [][]int64{
		{1, 2},
		{2, 1},
		{1, 1},
		{1, 2},
		{2, 1},
		{1, 1},
		{1, 2},
		{2, 1},
		{1, 1},
	}

	s.comparable = []Comparable{
		LESS_THAN,
		LESS_THAN,
		LESS_THAN,
		EQUAL,
		EQUAL,
		EQUAL,
		GREATER_THAN,
		GREATER_THAN,
		GREATER_THAN,
	}

	s.expected = []bool{
		true,
		false,
		false,
		false,
		false,
		true,
		false,
		true,
		false,
	}
}

func (s *SuiteComparable) TearDownTest() {
	fmt.Println("Ending Test")
}

func Test_ComparableSuite(t *testing.T) {
	suite.Run(t, new(SuiteComparable))
}
