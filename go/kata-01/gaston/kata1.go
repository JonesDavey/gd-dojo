package main

import (
	"fmt"
)

func binarySearch(target int, list []int) int {
	return getIndex(target, list, 0, len(list))
}

func getIndex(target int, list []int, start int, end int) int {
	midpoint := (start + end) / 2

	if start >= end {
		return -1 // Element is not in list
	}

	if list[midpoint] == target {
		return midpoint
	} else if list[midpoint] < target {
		// We need to look in the right half of the list
		return getIndex(target, list, midpoint+1, end)
	} else {
		// Look in the left half
		return getIndex(target, list, start, midpoint)
	}
}

func applyTest(expected, searchValue int, inputCollection []int) bool {
	actual := binarySearch(searchValue, inputCollection)
	return expected == actual
}

func main() {
	fmt.Println("test case: 0", applyTest(-1, 3, []int{}))
	fmt.Println("test case: 1", applyTest(-1, 3, []int{1}))
	fmt.Println("test case: 2", applyTest(0, 1, []int{1}))

	fmt.Println("test case: 3", applyTest(0, 1, []int{1, 3, 5}))
	fmt.Println("test case: 4", applyTest(1, 3, []int{1, 3, 5}))
	fmt.Println("test case: 5", applyTest(2, 5, []int{1, 3, 5}))
	fmt.Println("test case: 6", applyTest(-1, 0, []int{1, 3, 5}))
	fmt.Println("test case: 7", applyTest(-1, 2, []int{1, 3, 5}))
	fmt.Println("test case: 8", applyTest(-1, 4, []int{1, 3, 5}))
	fmt.Println("test case: 9", applyTest(-1, 6, []int{1, 3, 5}))

	fmt.Println("test case 1: ", applyTest(0, 1, []int{1, 3, 5, 7}))
	fmt.Println("test case 2: ", applyTest(1, 3, []int{1, 3, 5, 7}))
	fmt.Println("test case 3: ", applyTest(2, 5, []int{1, 3, 5, 7}))
	fmt.Println("test case 4: ", applyTest(3, 7, []int{1, 3, 5, 7}))
	fmt.Println("test case 5: ", applyTest(-1, 0, []int{1, 3, 5, 7}))
	fmt.Println("test case 6: ", applyTest(-1, 2, []int{1, 3, 5, 7}))
	fmt.Println("test case 7: ", applyTest(-1, 4, []int{1, 3, 5, 7}))
	fmt.Println("test case 9: ", applyTest(-1, 6, []int{1, 3, 5, 7}))
	fmt.Println("test case 10: ", applyTest(-1, 8, []int{1, 3, 5, 7}))
}
