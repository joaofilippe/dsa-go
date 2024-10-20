package binarytrees

import (
	"testing"
)

func TestNewBinaryTree(t *testing.T) {
	rootValue := 10
	tree := NewBinaryTree(rootValue)

	if tree == nil {
		t.Fatalf("Expected non-nil tree, got nil")
	}

	if tree.Root == nil {
		t.Fatalf("Expected non-nil root, got nil")
	}

	if tree.Root.Value != rootValue {
		t.Errorf("Expected root value %d, got %d", rootValue, tree.Root.Value)
	}

	if tree.Root.Left != nil {
		t.Errorf("Expected root left child to be nil, got non-nil")
	}

	if tree.Root.Right != nil {
		t.Errorf("Expected root right child to be nil, got non-nil")
	}
}
