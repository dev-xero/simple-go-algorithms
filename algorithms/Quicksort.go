package algorithms

func QuickSort(array []int) {
	if len(array) <= 1 {
		return
	}

	pivotIndex := partition(array)

	QuickSort(array[:pivotIndex])
	QuickSort(array[pivotIndex+1:])
}

func partition(array []int) int {
	pivotIndex := len(array) - 1
	pivot := array[pivotIndex]
	left := 0
	right := pivotIndex - 1

	for {
		for left <= right && array[left] < pivot {
			left++
		}

		for left <= right && array[right] > pivot {
			right--
		}

		if left <= right {
			array[left], array[right] = array[right], array[left]
			left++
			right--
		} else {
			break
		}
	}

	array[left], array[pivotIndex] = array[pivotIndex], array[left]

	return left
}
