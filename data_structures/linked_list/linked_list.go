package linkedlist

// NewLinkedListFromSlice returns a linked list of nodes from an array of integers
func NewLinkedListFromSlice(values []int) LinkedList {
	temp := new(Node)
	for i := 0; i < len(values); i++ {
		node := &Node{Value: values[i]}

		if temp == nil {
			temp = node
			continue
		}

		curr := temp
		for curr.Next != nil {
			curr = curr.Next
		}

		curr.Next = node
	}

	return LinkedList{Node: temp.Next}
}

// AddAtStart adds a node in
func (l *LinkedList) InsertAtStart(value int) {
	node := &Node{value, l.Node}
	l.Node = node
}
