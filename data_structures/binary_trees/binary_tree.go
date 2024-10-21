package binarytrees

type BinaryTree struct {
	Root *Node
}

func NewBinaryTree(rootValue int) *BinaryTree {
	return &BinaryTree{Root: &Node{Value: rootValue}}
}

func (bt *BinaryTree) Insert(value int) {
	if bt.Root == nil {
		bt.Root = &Node{Value: value}
		return
	}
	insert(bt.Root, value)
}

func insert(leaf *Node, value int) {
	queue := []*Node{leaf}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.Left == nil {
			current.Left = &Node{Value: value}
			return
		} else {
			queue = append(queue, current.Left)
		}

		if current.Right == nil {
			current.Right = &Node{Value: value}
			return
		} else {
			queue = append(queue, current.Right)
		}
	}
}
