package main

import (
	"fmt"
	"github.com/dev-xero/SimpleGoAlgorithms/algorithms"
)

func main() {
	test_slice := []int{1, 3, 5, 7, 9, 11, 13, 15, 17}
	key := 4
	index := algorithms.BinarySearch(key, test_slice) // -1 since 4 is absent

	gcdOfNums := algorithms.GCD(86, 32)
	factOfNum := algorithms.Fact(5) // 5! -> 120

	fmt.Printf("The index of %d in the slice is: %dn\n", key, index)
	fmt.Printf("The greatest common divisor of 86 and 32 is: %d\n", gcdOfNums)
	fmt.Printf("5! is: %d", factOfNum)
}
