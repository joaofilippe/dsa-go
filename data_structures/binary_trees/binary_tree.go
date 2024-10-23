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

func (bt *BinaryTree) PreOrder() []int {
	if bt.Root == nil {
		return []int{}
	}

	result := new([]int)	
	*result = append(*result, bt.Root.Value)
	preOrder(bt.Root.Left, result)
	preOrder(bt.Root.Right, result)

	return *result
}

func preOrder(node *Node, result *[]int) {
	if node == nil {
		return
	}

	*result = append(*result, node.Value)
	preOrder(node.Left, result)
	preOrder(node.Right, result)
}

func (bt *BinaryTree) InOrder() []int {
	if bt.Root == nil {
		return []int{}
	}

	result := new([]int)
	inOrder(bt.Root.Left, result)
	*result = append(*result, bt.Root.Value)
	inOrder(bt.Root.Right, result)

	return *result
}

func inOrder(node *Node, result *[]int) {
	if node == nil {
		return
	}

	inOrder(node.Left, result)
	*result = append(*result, node.Value)
	inOrder(node.Right, result)
}

func (bt *BinaryTree) PostOrder() []int {
	if bt.Root == nil {
		return []int{}
	}

	result := new([]int)
	postOrder(bt.Root.Left, result)
	postOrder(bt.Root.Right, result)
	*result = append(*result, bt.Root.Value)

	return *result
}

func postOrder(node *Node, result *[]int) {
	if node == nil {
		return
	}

	postOrder(node.Left, result)
	postOrder(node.Right, result)
	*result = append(*result, node.Value)
}

