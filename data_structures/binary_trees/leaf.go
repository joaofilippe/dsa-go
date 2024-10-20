package binarytrees

type Leaf struct {
	Value int
	Left  *Leaf
	Right *Leaf
}

func NewLeaf(value int) *Leaf {
	return &Leaf{Value: value}
}

func (l *Leaf) Insert(value int) {
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