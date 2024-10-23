package binarytrees

import "testing"

func TestInsertNode(t *testing.T) {
	tests := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	}

	for _, test := range tests {
		leaf := NewNode(0)
		leaf.Insert(test)

		if leaf.Value != 0 {
			t.Errorf("Expected root value to be 0, got %d", leaf.Value)
		}
	}
}



func TestHeight(t *testing.T) {
	tests := []struct {
		values   []int
		expected int
	}{
		{[]int{}, 1},
		{[]int{1}, 2},
		{[]int{1, 2}, 3},
		{[]int{1, 2, 3}, 4},
		{[]int{3, 2, 1}, 4},
		{[]int{1, 2, 3, 4}, 5},
		{[]int{4, 3, 2, 1}, 5},
		{[]int{1, 2, 3, 4, 5}, 6},
	}

	for _, test := range tests {
		root := NewNode(0)
		for _, value := range test.values {
			root.Insert(value)
		}

		if height := root.Height(); height != test.expected {
			t.Errorf("For values %v, expected height %d, got %d", test.values, test.expected, height)
		}
	}
}

