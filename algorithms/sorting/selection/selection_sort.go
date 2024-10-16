package sorting

// SelectionSortByMaxIndex is a func to sort an integer list by selection method with the max index
func SelectionSortByMaxIndex(arr []int64) {
	size := len(arr)
	for i := 0; i < size; i++ {
		maxIdx := 0
		sizeI := size - 1 - i
		for j := 1; j <= sizeI; j++ {
			if arr[j] > arr[maxIdx] {
				maxIdx = j
			}
		}
		arr[sizeI], arr[maxIdx] = arr[maxIdx], arr[sizeI]
	}

}

// SelectionSortByMinIndex is a func to sort an integer list by selection method with the max index
func SelectionSortByMinIndex(arr []int64) {
	size := len(arr)
	for i := 0; i < size; i++ {
		minIdx := 0
		sizeI := size - 1 - i
		for j := 1; j <= sizeI; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[sizeI], arr[minIdx] = arr[minIdx], arr[sizeI]
	}

}
