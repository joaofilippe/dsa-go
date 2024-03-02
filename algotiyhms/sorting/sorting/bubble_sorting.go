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
