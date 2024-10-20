package binarytrees

type Root struct {
	Value int
	Left  *Leaf
	Right *Leaf
}

type BinaryTree struct {
	Root *Root
}

func NewBinaryTree(rootValue int) *BinaryTree {
	return &BinaryTree{Root: &Root{Value: rootValue}}
}
