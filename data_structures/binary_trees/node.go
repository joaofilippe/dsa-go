package binarytrees

type Node struct {
	Key   int
	Value int
	Left  *Node
	Right *Node
}

func NewNode(key, value int) *Node {
	return &Node{
		Key:   key,
		Value: value,
	}
}

func (n *Node) Insert(key, value int) {
	if value < n.Value {
		if n.Left == nil {
			n.Left = NewNode(key, value)
		} else {
			n.Left.Insert(key, value)
		}
	} else {
		if n.Right == nil {
			n.Right = NewNode(key, value)
		} else {
			n.Right.Insert(key, value)
		}
	}
}

func (n *Node) Height() int {
	if n == nil {
		return 0
	}

	leftHeight := n.Left.Height()
	rightHeight := n.Right.Height()

	if leftHeight > rightHeight {
		return leftHeight + 1
	}

	return rightHeight + 1
}
