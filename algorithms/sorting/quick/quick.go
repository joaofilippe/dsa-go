package sorting

func QuickSorting(input []int) []int {
	temp := make([]int, len(input))

	copy(temp, input)

	return quickSort(temp, 0, len(temp)-1)
}

func quickSort(input []int, start, end int) []int {
	if start < end {
		input, pivot := partition(input, start, end)
		quickSort(input, start, pivot-1)
		quickSort(input, pivot+1, end)
	}

	return input
}

func partition(input []int, start, end int) ([]int, int) {
	pivot := input[end]
	i := start
	for j := start; j < end; j++ {
		if input[j] <= pivot {
			input[i], input[j] = input[j], input[i]
			i++
		}
	}

	input[i], input[end] = input[end], input[i]
	return input, i
}
