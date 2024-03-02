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

// InsertAtStart adds a node in the begging of the list
func (l *LinkedList) InsertAtStart(value int) {
	node := &Node{value, l.Node}
	l.Node = node
}

// InsertAtEnd adds a node in the end of the list
func (l *LinkedList) InsertAtEnd(value int) {

}
