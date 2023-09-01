package algorithms

/*
Simple recursive function to compute the floor of log of a number
to a base
*/
func FloorOfLog(num int, base int, count int) int {
	if num <= 1 {
		return count
	}

	return FloorOfLog(num/base, base, count+1)
}
