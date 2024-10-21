package binarytrees

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func NewLeaf(value int) *Node {
	return &Node{Value: value}
}

func (l *Node) Insert(value int) {
	if value < l.Value {
		if l.Left == nil {
			l.Left = NewLeaf(value)
		} else {
			l.Left.Insert(value)
		}
	} else {
		if l.Right == nil {
			l.Right = NewLeaf(value)
		} else {
			l.Right.Insert(value)
		}
	}
}
