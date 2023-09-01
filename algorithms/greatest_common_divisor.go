package algorithms

/*
	Euclid's algorithm for determining the greatest common divisor of two
	numbers
*/
func GCD(dividend int, divisor int) int {
	remainder := dividend % divisor

	if remainder == 0 {
		return divisor
	}

	return GCD(divisor, remainder)
}