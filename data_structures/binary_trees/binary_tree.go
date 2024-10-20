package binarytrees

type BinaryTree struct {
	Root *Leaf
}

func NewBinaryTree(rootValue int) *BinaryTree {
	return &BinaryTree{Root: &Leaf{Value: rootValue}}
}

func (b *BinaryTree) InsertLeaf(value int) {
	leaf := NewLeaf(value)
	if b.Root.Left == nil {
		b.Root.Left = leaf
	} else if b.Root.Right == nil {
		b.Root.Right = leaf
	} else {
		insertLeafRecursive(b.Root)
	}
}

func insertLeafRecursive(root *Leaf) {
	if root.Left == nil {
		root.Left = root
	} else if root.Right == nil {
		root.Right = root
	} else {
		insertLeafRecursive(root.Left)
	}
}
