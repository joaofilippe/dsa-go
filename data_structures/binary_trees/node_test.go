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
