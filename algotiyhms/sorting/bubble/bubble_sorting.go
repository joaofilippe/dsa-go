package sorting

import (
	"github.com/joaofilippe/dsa-go/enums"
)

// BubbleSorting is a func wich returns a sorted slice by bubble method
func BubbleSorting(slice []int64, comparable enums.Comparable) []int64{
	size := len(slice)
	for i := 0; i < size - 1; i++ {
		for j := 0; j < size - 1 - i; j++ {
			if comparable.Compare(slice[j], slice[j + 1]){
				slice[j], slice[j + 1] = slice[j + 1], slice[j]
			}
		}
	}

	return slice
}

/* 
1.	The outer loop represents the number of swaps that are done for the comparison of data.

2.	The inner loop is used for the comparison of data. In the first iteration,
	the largest value will be  moved to the end of the array, in the second iteration the
	second-largest value will be moved  before the largest value and so on.

3.	The greater() function is used for comparison, which means when the value of 
	the first  argument is greater than the value of the second argument then performs a swap.
	With this, we  are sorting in increasing order. If we have, the less() function in place
	of greater() then the array  will be sorted in decreasing order. 

	Jain, Hemant. Data Structures & Algorithms In Go (p. 49). Edição do Kindle. 
*/