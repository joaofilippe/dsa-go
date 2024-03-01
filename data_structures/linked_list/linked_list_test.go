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

		assert.True(t, checkLinkedList(linkedList,test), "the linked list values doesn't matche with slice values")
		assert.NotNil(t, linkedList, "the linked_list must no be nil")
	}
}

func Test_AddAtStart(t *testing.T){
	tests := []struct{
		slice []int
		newValue int
		newSlice []int
	}{
		{
			[]int{1,2,3,4,5,6},
			7,
			[]int{7,1,2,3,4,5,6},
		},
		{
			[]int{1,2,3,4,5,6},
			7,
			[]int{7,1,2,3,4,5,6},
		},
		{
			[]int{1,2,3,4,5,6},
			7,
			[]int{7,1,2,3,4,5,6},
		},
		{
			[]int{1,2,3,4,5,6},
			7,
			[]int{7,1,2,3,4,5,6},
		},
		{
			[]int{1,2,3,4,5,6},
			7,
			[]int{7,1,2,3,4,5,6},
		},
		{
			[]int{1,2,3,4,5,6},
			7,
			[]int{7,1,2,3,4,5,6},
		},
	}

	for _, test := range tests {
		linkedList := NewLinkedListFromSlice(test.slice)
		linkedList.InsertAtStart(test.newValue)

		result := checkLinkedList(linkedList, test.newSlice)
		fmt.Println(result)
	}
}

func checkLinkedList(linkedList LinkedList, slice []int) bool {
	i := 0
	for curr := linkedList.Node; curr.Next != nil; curr = curr.Next{
		if curr.Value != slice[i] {
			return false
		}
		i++
	}

	return true
}
