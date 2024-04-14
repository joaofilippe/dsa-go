package sorting

// SelectionSort is a func to sort an integer list by selection method
func SelectionSortByMaxIndex(arr []int64) {
	size := len(arr)
	for i := 0; i < size; i++ {
		maxIdx := 0
		sizeI := size - 1 - i
		for j := 1; j <= sizeI; j++ {
			arrj := arr[j]
			arrmaxIdx := arr[maxIdx]
			if arrj > arrmaxIdx {
				maxIdx = j
			}
		}
		arr[sizeI], arr[maxIdx] = arr[maxIdx], arr[sizeI]
	}

}
