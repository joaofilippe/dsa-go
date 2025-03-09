package linkedlist

type LinkedList struct {
	Node *Node
}

type Node struct {
	Index int
	Value int
	Next *Node
}
