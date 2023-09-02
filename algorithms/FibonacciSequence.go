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

func FastFib(num int) string {
	var fibCache = map[int]int{}
	sequence := []string{"1", "1"}

	if num <= 0 {
		return ""
	} else if num == 1 {
		return "1"
	} else if num == 2 {
		return "1 1"
	}

	fibCache[1], fibCache[2] = 1, 1

	for i := 3; i <= num; i++ {
		fib := getFib(i, fibCache)
		sequence = append(sequence, strconv.Itoa(fib))
	}

	return strings.Join(sequence, " ")
}

func getFib(n int, cache map[int]int) int {
	if val, ok := cache[n]; ok {
		return val
	}

	cache[n] = getFib(n-1, cache) + getFib(n-2, cache)

	return cache[n]
}
