package linkedlist

import (
	"bytes"
	"io"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func checkLinkedList(linkedList LinkedList, slice []int) bool {
	i := 0
	for curr := linkedList.Node; curr.Next != nil; curr = curr.Next {
		if curr.Value != slice[i] {
			return false
		}
		i++
	}

	return true
}

func Test_NewLinkedListFromSlice(t *testing.T) {
	tests := [][]int{
		{1, 2, 3, 4, 5, 6},
		{6, 5, 4, 3, 2, 1},
		{20, 35, 89, 90, 121, 144},
	}

	for _, test := range tests {
		linkedList := NewLinkedListFromSlice(test)

		assert.True(t, checkLinkedList(linkedList, test), "the linked list values doesn't matche with slice values")
		assert.NotNil(t, linkedList, "the linked_list must no be nil")
	}
}

func Test_InsertAtStart(t *testing.T) {
	tests := []struct {
		slice    []int
		newValue int
		newSlice []int
	}{
		{
			[]int{1, 2, 3, 4, 5, 6},
			7,
			[]int{7, 1, 2, 3, 4, 5, 6},
		},
		{
			[]int{1, 2, 3, 4, 5, 6},
			8,
			[]int{8, 1, 2, 3, 4, 5, 6},
		},
		{
			[]int{2, 3, 4, 5, 6},
			1,
			[]int{1, 2, 3, 4, 5, 6},
		},
	}

	for _, test := range tests {
		linkedList := NewLinkedListFromSlice(test.slice)
		linkedList.InsertAtStart(test.newValue)

		result := checkLinkedList(linkedList, test.newSlice)
		assert.True(t, result, "lists doesn't match")
	}
}

func Test_InsertAtEnd(t *testing.T) {
	tests := []struct {
		slice    []int
		newValue int
		newSlice []int
	}{
		{
			[]int{1, 2, 3, 4, 5, 6},
			7,
			[]int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			[]int{1, 2, 3, 4, 5, 6},
			8,
			[]int{1, 2, 3, 4, 5, 6, 8},
		},
		{
			[]int{2, 3, 4, 5, 6},
			1,
			[]int{2, 3, 4, 5, 6, 1},
		},
	}

	for _, test := range tests {
		linkedList := NewLinkedListFromSlice(test.slice)
		linkedList.InsertAtEnd(test.newValue)

		result := checkLinkedList(linkedList, test.newSlice)
		assert.True(t, result, "lists doesn't match")
	}
}

func Test_InsertAfter(t *testing.T) {
	tests := []struct {
		slice    []int
		newValue int
		index    int
		newSlice []int
	}{
		{
			[]int{1, 2, 3, 4, 5, 6},
			7,
			3,
			[]int{1, 2, 3, 4, 7, 5, 6},
		},
		{
			[]int{1, 2, 3, 4, 5, 6},
			8,
			4,
			[]int{1, 2, 3, 4, 5, 8, 6},
		},
		{
			[]int{2, 3, 4, 5, 6},
			1,
			0,
			[]int{2, 1, 3, 4, 5, 6, 1},
		},
		{
			[]int{2, 3, 4, 5, 6},
			1,
			-1,
			[]int{1, 2, 3, 4, 5, 6, 1},
		},
	}

	for _, test := range tests {
		linkedList := NewLinkedListFromSlice(test.slice)
		linkedList.InsertAfter(test.index, test.newValue)

		result := checkLinkedList(linkedList, test.newSlice)
		assert.True(t, result, "lists doesn't match")
	}
}

func Test_InsertAtPosition(t *testing.T) {
	tests := []struct {
		slice    []int
		position int
		newValue int
		newSlice []int
	}{
		{
			[]int{1, 2, 3, 4, 5, 6},
			0,
			7,
			[]int{7, 1, 2, 3, 4, 5, 6},
		},
		{
			[]int{1, 2, 3, 4, 5, 6},
			3,
			8,
			[]int{1, 2, 3, 8, 4, 5, 6},
		},
		{
			[]int{2, 3, 4, 5, 6},
			5,
			1,
			[]int{2, 3, 4, 5, 6, 1},
		},
		{
			[]int{2, 3, 4, 5, 6},
			2,
			9,
			[]int{2, 3, 9, 4, 5, 6},
		},
	}

	for _, test := range tests {
		linkedList := NewLinkedListFromSlice(test.slice)
		linkedList.InsertAtPosition(test.position, test.newValue)

		result := checkLinkedList(linkedList, test.newSlice)
		assert.True(t, result, "lists doesn't match")
	}
}
func Test_DeleteAtPosition(t *testing.T) {
	tests := []struct {
		slice    []int
		position int
		newSlice []int
	}{
		{
			[]int{1, 2, 3, 4, 5, 6},
			0,
			[]int{2, 3, 4, 5, 6},
		},
		{
			[]int{1, 2, 3, 4, 5, 6},
			3,
			[]int{1, 2, 3, 5, 6},
		},
		{
			[]int{2, 3, 4, 5, 6},
			4,
			[]int{2, 3, 4, 5},
		},
		{
			[]int{2, 3, 4, 5, 6},
			2,
			[]int{2, 3, 5, 6},
		},
	}

	for _, test := range tests {
		linkedList := NewLinkedListFromSlice(test.slice)
		linkedList.DeleteAtPosition(test.position)

		result := checkLinkedList(linkedList, test.newSlice)
		assert.True(t, result, "lists doesn't match")
	}
}
func Test_DeleteAtEnd(t *testing.T) {
	tests := []struct {
		slice    []int
		newSlice []int
	}{
		{
			[]int{1, 2, 3, 4, 5, 6},
			[]int{1, 2, 3, 4, 5},
		},
		{
			[]int{1, 2, 3, 4, 5},
			[]int{1, 2, 3, 4},
		},
		{
			[]int{2, 3, 4, 5, 6},
			[]int{2, 3, 4, 5},
		},
		{
			[]int{1},
			[]int{},
		},
	}

	for _, test := range tests {
		linkedList := NewLinkedListFromSlice(test.slice)
		linkedList.DeleteAtEnd()

		result := checkLinkedList(linkedList, test.newSlice)
		assert.True(t, result, "lists doesn't match")
	}
}
func Test_DeleteAtStart(t *testing.T) {
	tests := []struct {
		slice    []int
		newSlice []int
	}{
		{
			[]int{1, 2, 3, 4, 5, 6},
			[]int{2, 3, 4, 5, 6},
		},
		{
			[]int{1, 2, 3, 4, 5},
			[]int{2, 3, 4, 5},
		},
		{
			[]int{2, 3, 4, 5, 6},
			[]int{3, 4, 5, 6},
		},
		{
			[]int{1},
			[]int{},
		},
	}

	for _, test := range tests {
		linkedList := NewLinkedListFromSlice(test.slice)
		linkedList.DeleteAtStart()

		result := checkLinkedList(linkedList, test.newSlice)
		assert.True(t, result, "lists doesn't match")
	}
}
func Test_Reverse(t *testing.T) {
	tests := []struct {
		slice    []int
		expected []int
	}{
		{
			[]int{1, 2, 3, 4, 5, 6},
			[]int{6, 5, 4, 3, 2, 1},
		},
		{
			[]int{1, 2, 3, 4, 5},
			[]int{5, 4, 3, 2, 1},
		},
		{
			[]int{2, 3, 4, 5, 6},
			[]int{6, 5, 4, 3, 2},
		},
		{
			[]int{1},
			[]int{1},
		},
		{
			[]int{},
			[]int{},
		},
	}

	for _, test := range tests {
		linkedList := NewLinkedListFromSlice(test.slice)
		linkedList.Reverse()

		result := checkLinkedList(linkedList, test.expected)
		assert.True(t, result, "lists doesn't match")
	}
}

func Test_Print(t *testing.T) {
	tests := []struct {
		slice    []int
		expected string
	}{
		{
			[]int{1, 2, 3, 4, 5, 6},
			"123456",
		},
		{
			[]int{1, 2, 3, 4, 5},
			"12345",
		},
		{
			[]int{2, 3, 4, 5, 6},
			"23456",
		},
		{
			[]int{1},
			"1",
		},
		{
			[]int{},
			"",
		},
	}

	for _, test := range tests {
		linkedList := NewLinkedListFromSlice(test.slice)
		output := captureOutput(func() {
			err := linkedList.Print()
			assert.Nil(t, err)
		})
		assert.Equal(t, test.expected, output)
	}
}

func Test_ToSlice(t *testing.T) {
	tests := []struct {
		slice    []int
		expected []int
	}{
		{
			[]int{1, 2, 3, 4, 5, 6},
			[]int{1, 2, 3, 4, 5, 6},
		},
		{
			[]int{1, 2, 3, 4, 5},
			[]int{1, 2, 3, 4, 5},
		},
		{
			[]int{2, 3, 4, 5, 6},
			[]int{2, 3, 4, 5, 6},
		},
		{
			[]int{1},
			[]int{1},
		},
		{
			[]int{},
			[]int{},
		},
	}

	for _, test := range tests {
		linkedList := NewLinkedListFromSlice(test.slice)
		result := ToSlice(linkedList)
		assert.Equal(t, test.expected, result)
	}
}

func Test_Length(t *testing.T) {
	tests := []struct {
		slice    []int
		expected int
	}{
		{
			[]int{1, 2, 3, 4, 5, 6},
			6,
		},
		{
			[]int{1, 2, 3, 4, 5},
			5,
		},
		{
			[]int{2, 3, 4, 5, 6},
			5,
		},
		{
			[]int{1},
			1,
		},
		{
			[]int{},
			0,
		},
	}

	for _, test := range tests {
		linkedList := NewLinkedListFromSlice(test.slice)
		result := linkedList.Length()
		assert.Equal(t, test.expected, result)
	}
}

// Helper function to capture output
func captureOutput(f func()) string {
	var buf bytes.Buffer
	writer := io.Writer(&buf)
	log.SetOutput(writer)
	defer func() {
		log.SetOutput(os.Stdout)
	}()
	f()
	return buf.String()
}
