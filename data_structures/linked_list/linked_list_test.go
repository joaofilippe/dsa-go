package linkedlist

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewLinkedListFromSlice(t *testing.T) {
	tests := [][]int{
		{1, 2, 3, 4, 5, 6},
		{6, 5, 4, 3, 2, 1},
		{20, 35, 89, 90, 121, 144},
	}

	for _, test := range tests {
		linkedList := NewLinkedListFromSlice(test)

		assert.NotNil(t, linkedList, "the linked_list must no be nil")
	}
}

func Test_Tests(t *testing.T){
	l := &LinkedList{
		Node: &Node{
			2,
			&Node{
				3,
				&Node{
					3,
					&Node{
						4,
						nil,
					},
				},
			},
		},
	}

	for curr := l.Node; curr != nil; curr = curr.Next {
		fmt.Println(curr.Value)
	}
}