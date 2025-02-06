package linkedlist

import "fmt"

// NewLinkedListFromSlice returns a linked list of nodes from an array of integers
func NewLinkedListFromSlice(values []int) LinkedList {
	temp := new(Node)
	for i := 0; i < len(values); i++ {
		node := &Node{Value: values[i]}

		curr := temp
		for curr.Next != nil {
			curr = curr.Next
		}

		curr.Next = node
	}

	return LinkedList{Node: temp.Next}
}

func (l *LinkedList) Print() error {
	if l.Node == nil {
		return fmt.Errorf("list is empty")
	}

	current := l.Node
	for current != nil {
		fmt.Printf("%v", current.Value)
		current = current.Next
	}

	return nil
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

// Insert adds a node after the node with the index informed
func (l *LinkedList) Insert(index, value int) {
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

// ToSlice converts a LinkedList to a slice of integers.
// It iterates through each node in the linked list, appending the value of each node to a slice.
// The function returns the resulting slice of integers.
func ToSlice(l LinkedList) []int {
	slice := []int{}
	curr := l.Node
	for curr != nil {
		slice = append(slice, curr.Value)
		curr = curr.Next
	}

	return slice
}

// DeleteAtStart removes the first node from the linked list.
// After the operation, the head of the list will be the next node.
// If the list is empty, this operation has no effect.
func (l *LinkedList) DeleteAtStart() {
	l.Node = l.Node.Next
}

// DeleteAtEnd removes the last node from the linked list.
// It traverses the list to find the second-to-last node and sets its Next pointer to nil,
// effectively removing the last node from the list.
func (l *LinkedList) DeleteAtEnd() {
	curr := l.Node
	for curr.Next.Next != nil {
		curr = curr.Next
	}

	curr.Next = nil
}

// DeleteAtPosition removes the node at the specified position in the linked list.
// If the position is 0, it deletes the node at the start of the list.
// If the position is greater than 0, it traverses the list to find the node at the given position
// and removes it by updating the pointers of the previous node.
//
// Parameters:
//   position (int): The zero-based index of the node to be deleted.
//
// Note:
//   This function assumes that the position is within the bounds of the list.
func (l *LinkedList) DeleteAtPosition(position int) {
	if position == 0 {
		l.DeleteAtStart()
		return
	}

	curr := l.Node
	for i := 0; i < position-1; i++ {
		curr = curr.Next
	}

	curr.Next = curr.Next.Next
}

// Reverse reverses the order of the nodes in the linked list.
// It iterates through the list, reversing the direction of the links
// between the nodes, so that the head of the list becomes the tail,
// and the tail becomes the head.
func (l *LinkedList) Reverse() {
	var prev *Node
	curr := l.Node
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	l.Node = prev
}

func (l *LinkedList) Length() int {
	count := 0
	curr := l.Node
	for curr != nil {
		count++
		curr = curr.Next
	}

	return count
}
