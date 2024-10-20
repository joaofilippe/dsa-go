package binarytrees

type BinaryTree struct {
	Root *Leaf
}

func NewBinaryTree(rootValue int) *BinaryTree {
	return &BinaryTree{Root: &Leaf{Value: rootValue}}
}



func (bt *BinaryTree) Insert(value int) {
	if bt.Root == nil {
		bt.Root = &Leaf{Value: value}
		return
	}
	insert(bt.Root, value)
}

func insert(leaf *Leaf, value int) {
	queue := []*Leaf{leaf}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.Left == nil {
			current.Left = &Leaf{Value: value}
			return
		} else {
			queue = append(queue, current.Left)
		}

		if current.Right == nil {
			current.Right = &Leaf{Value: value}
			return
		} else {
			queue = append(queue, current.Right)
		}
	}
}