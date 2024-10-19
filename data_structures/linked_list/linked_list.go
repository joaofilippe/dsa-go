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
	curr := l.Node.Next
	for curr.Next != nil {
		curr = curr.Next
	}

	curr.Next = &Node{Value: value, Next: nil}
}

// InsertAfter adds a node after the node with the index informed
func (l *LinkedList) InsertAfter(index, value int) {
	if index == -1 {
		l.InsertAtStart(value)
		return
	}

	curr := l.Node
	for i := 0; i < index; i++ {
		curr = curr.Next
	}

	node := &Node{Value: value, Next: curr.Next}
	curr.Next = node
}

 // InsertAtPosition inserts a new node with the given value at the specified position in the linked list.
// If the position is 0, it inserts the node at the start of the list.
// If the position is greater than the length of the list, it will panic due to dereferencing a nil pointer.
//
// Parameters:
//   - position: The zero-based index at which the new node should be inserted.
//   - value: The value to be stored in the new node.
func (l *LinkedList) InsertAtPosition(position int, value int) {
	if position == 0 {
		l.InsertAtStart(value)
		return
	}

	curr := l.Node
	for i := 0; i < position-1; i++ {
		curr = curr.Next
	}

	node := &Node{Value: value, Next: curr.Next}
	curr.Next = node
}

func ToSlice(l LinkedList) []int {
	slice := []int{}
	curr := l.Node
	for curr != nil {
		slice = append(slice, curr.Value)
		curr = curr.Next
	}

	return slice
}