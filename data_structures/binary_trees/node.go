package binarytrees

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func NewNode(value int) *Node {
	return &Node{Value: value}
}

func (n *Node) Insert(value int) {
	if value < n.Value {
		if n.Left == nil {
			n.Left = NewNode(value)
		} else {
			n.Left.Insert(value)
		}
	} else {
		if n.Right == nil {
			n.Right = NewNode(value)
		} else {
			n.Right.Insert(value)
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