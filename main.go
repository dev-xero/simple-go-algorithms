package main

import (
	"fmt"
	binary "github.com/dev-xero/SimpleGoAlgorithms/algorithms"
)

func main() {
	test_slice := []int{1, 3, 5, 7, 9, 11, 13, 15, 17}
	key := 4

	index := binary.Search(key, test_slice)

	fmt.Printf("The index of %d in the slice is: %d", key, index)
}
