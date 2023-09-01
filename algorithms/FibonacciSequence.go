package algorithms

import (
	"strconv"
	"strings"
)

/*
Algorithm to return a fibonacci sequence
up until a certain number of terms
*/
func Fib(num int) string {
	cache := []int{1, 1}
	sequence := []string{"1", "1"}

	for i := 2; i < num; i++ {
		cache = append(cache, cache[i-1]+cache[i-2])
		sequence = append(sequence, strconv.Itoa(cache[i]))
	}

	return strings.Join(sequence, " ")
}
