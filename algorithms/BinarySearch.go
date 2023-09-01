package algorithms

/*
Binary Search
searches for a key in an ordered array in log-time and returns its index
returns -1 if the array doesn't contain the key
*/
func BinarySearch(key int, array []int) int {
	lowerBound := 0
	upperBound := len(array) - 1

	for lowerBound < upperBound {

		midPoint := lowerBound + ((upperBound - lowerBound) / 2)

		if array[midPoint] == key {
			return midPoint
		} else if array[midPoint] > key {
			upperBound = midPoint - 1
		} else {
			lowerBound = midPoint + 1
		}
	}

	return -1
}
