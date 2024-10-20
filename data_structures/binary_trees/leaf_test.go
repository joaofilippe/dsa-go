package binarytrees

import "testing"

func TestInsertLeaf(t *testing.T) {
	tests := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	}

	for _, test := range tests {
		leaf := NewLeaf(0)
		leaf.Insert(test)

		if leaf.Value != 0 {
			t.Errorf("Expected root value to be 0, got %d", leaf.Value)
		}
	}
}