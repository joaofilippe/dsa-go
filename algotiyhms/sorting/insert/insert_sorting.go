package sorting

import (
	"github.com/joaofilippe/dsa-go/enums"
)

// InsertionSort sorts a slice with insertion sort method
func InsertionSort(slice []int64, comparable enums.Comparable) {
	// needs start where could be a previous element to compare
	var i,j int
	for i = 1; i < len(slice); i++ {
		selected := slice[i]
		for j = i; j > 0 &&
			comparable.Compare(slice[j-1], int64(selected)); j-- {
			slice[j] = slice[j-1]
		}
		
		slice[j] = selected
	}
}

/*
1. 	The outer loop is used to choose the value to be inserted into the sorted array on the left.

2. 	The chosen value we want to insert is saved in a temp variable.

3.	The inner loop is doing the comparison using the greater() function.
	The values are shifted to  the right until we find the proper position of the temp value,
	for which this iteration is performed.

4. 	Finally, temp's value is placed in its proper position. In each iteration of the outer loop,
	the  length of the sorted array increases by one. When we exit the outer loop, the array is sorted.

	Jain, Hemant. Data Structures & Algorithms In Go (p. 51). Edição do Kindle.
*/
