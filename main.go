package main

import (
	"fmt"
	"github.com/dev-xero/SimpleGoAlgorithms/algorithms"
)

func main() {
	testSlice := []int{1, 3, 5, 7, 9, 11, 13, 15, 17}
	key := 4
	index := algorithms.BinarySearch(key, testSlice) // -1 since 4 is absent

	gcdOfNums := algorithms.GCD(86, 32)
	factOfNum := algorithms.Fact(5) // 5! -> 120
	floorOfLogOfNum := algorithms.FloorOfLog(7, 2, 0)
	testFib := algorithms.Fib(10)

	fmt.Printf("\nFibonacci Sequence with ten terms: %s\n", testFib)

	fmt.Printf("The index of %d in the slice is: %d\n", key, index)
	fmt.Printf("The greatest common divisor of 86 and 32 is: %d\n", gcdOfNums)
	fmt.Printf("5! is: %d\n", factOfNum)
	fmt.Printf("The floor of log of 7 to base 2 is: %d", floorOfLogOfNum)
}
