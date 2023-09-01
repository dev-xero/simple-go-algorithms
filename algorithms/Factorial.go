package algorithms

/*
	Factorial function (!)
*/
func Fact(num int) int {
	if num == 1 {
		return num
	}

	return num * Fact(num-1)
}
