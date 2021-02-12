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

func testChop(target int, expected int) {
	testArray := []int{1, 3, 4, 7, 13, 14, 15, 24}

	index := binarySearch(target, testArray[:])

	if index == expected {
		fmt.Println("Test passed.")
	} else {
		fmt.Printf("Test failed. Function returned %d, expected %d\n", index, expected)
	}
}

func main() {
	// Test edges
	testChop(1, 0)
	testChop(24, 7)

	// test a couple in the middle
	testChop(7, 3)
	testChop(14, 5)

	// Make sure it returns -1 if not in list
	testChop(8, -1)
}
