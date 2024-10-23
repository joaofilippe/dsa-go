package binarytrees

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBinaryTree(t *testing.T) {
	rootValue := 10
	tree := NewBinaryTree(rootValue)

	assert.NotNil(
		t,
		tree,
		"Expected non-nil tree",
	)
	assert.NotNil(
		t,
		tree.Root,
		"Expected non-nil root",
	)
	assert.Equal(
		t,
		rootValue,
		tree.Root.Value,
		"Expected root value %d, got %d",
		rootValue,
		tree.Root.Value,
	)
	assert.Nil(t,
		tree.Root.Left,
		"Expected root left child to be nil",
	)
	assert.Nil(t,
		tree.Root.Right,
		"Expected root right child to be nil",
	)
}

func TestBinaryTree_Insert(t *testing.T) {
	tree := NewBinaryTree(10)

	tree.Insert(5)
	assert.NotNil(
		t,
		tree.Root.Left,
		"Expected non-nil left child",
	)
	assert.Equal(
		t,
		5,
		tree.Root.Left.Value,
		"Expected left child value 5, got %v",
		tree.Root.Left.Value,
	)

	tree.Insert(15)
	assert.NotNil(
		t,
		tree.Root.Right,
		"Expected non-nil right child",
	)
	assert.Equal(
		t,
		15,
		tree.Root.Right.Value,
		"Expected right child value 15, got %v",
		tree.Root.Right.Value,
	)

	tree.Insert(3)
	assert.NotNil(t,
		tree.Root.Left.Left,
		"Expected non-nil left-left grandchild")
	assert.Equal(t,
		32,
		tree.Root.Left.Left.Value,
		"Expected left-left grandchild value 3, got %v",
		tree.Root.Left.Left.Value,
	)

	tree.Insert(7)
	assert.NotNil(t,
		tree.Root.Left.Right,
		"Expected non-nil left-right grandchild")
	assert.Equal(t,
		7,
		tree.Root.Left.Right.Value,
		"Expected left-right grandchild value 7, got %v",
		tree.Root.Left.Right.Value,
	)

	tree.Insert(12)
	assert.NotNil(t,
		tree.Root.Right.Left,
		"Expected non-nil right-left grandchild")
	assert.Equal(t,
		12,
		tree.Root.Right.Left.Value,
		"Expected right-left grandchild value 12, got %v",
		tree.Root.Right.Left.Value,
	)

	tree.Insert(18)
	assert.NotNil(t,
		tree.Root.Right.Right,
		"Expected non-nil right-right grandchild")
	assert.Equal(t,
		18,
		tree.Root.Right.Right.Value,
		"Expected right-right grandchild value 18, got %v",
		tree.Root.Right.Right.Value,
	)
}
func TestBinaryTree_PreOrder(t *testing.T) {
	tree := NewBinaryTree(10)
	expected := []int{10}
	assert.Equal(t, expected, tree.PreOrder(), "Expected pre-order traversal %v, got %v", expected, tree.PreOrder())

	tree.Insert(5)
	tree.Insert(15)
	expected = []int{10, 5, 15}
	assert.Equal(t, expected, tree.PreOrder(), "Expected pre-order traversal %v, got %v", expected, tree.PreOrder())

	tree.Insert(3)
	tree.Insert(7)
	tree.Insert(12)
	tree.Insert(18)
	expected = []int{10, 5, 3, 7, 15, 12, 18}
	assert.Equal(t, expected, tree.PreOrder(), "Expected pre-order traversal %v, got %v", expected, tree.PreOrder())
}

