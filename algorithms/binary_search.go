package algorithms

/*
	Binary Search
	searches for a key in an ordered array in log-time and returns its index
	returns -1 if the array doesn't contain the key
 */
func BinarySearch(key int, array []int) int {
	lower_bound := 0
	upper_bound := len(array) - 1

	for {
		if lower_bound >= upper_bound {
			break
		}

		mid_point := lower_bound + ((upper_bound - lower_bound) / 2)

		if array[mid_point] == key {
			return mid_point
		} else if array[mid_point] > key {
			upper_bound = mid_point - 1
		} else {
			lower_bound = mid_point + 1
		}
	}

	return -1
}
