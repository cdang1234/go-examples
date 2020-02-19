package search

import (
	"fmt"
	"math"
	"sort"
)

func BinarySearch() {
	target := 7
	arr := []int{-2, -1, 0, 5, 7, 7, 8, 19, 30, 56, 57, 98, 100}

	matches := sort.Search(len(arr), func(i int) bool {
		return arr[i] >= target
	})

	fmt.Println(matches)
}

// BinarySearchLessThan searches for first index with value less than target
func BinarySearchLessThan() {
	target := 7
	arr := []int{-2, -1, 0, 5, 7, 7, 8, 19, 30, 56, 57, 98, 100}

	matches := -1 + sort.Search(len(arr), func(i int) bool {
		return arr[i] >= target
	})

	fmt.Println(matches)
}

// BinarySearchGreaterThan searches for first index with value greater than target
func BinarySearchGreaterThan() {
	target := 7
	arr := []int{-2, -1, 0, 5, 7, 7, 8, 19, 30, 56, 57, 98, 100}

	matches := sort.Search(len(arr), func(i int) bool {
		return arr[i] > target
	})

	fmt.Println(matches)
}

// taken from capacity-to-ship-packages-within-d-days
func CustomBinarySearch() {
	min := math.MinInt32
	max := 0
	weights := []int{1, 2, 3, 4, 5, 6}
	D := 10

	works := func(i int, w []int) int { return 0 }

	for min <= max {
		curr := (max + min) / 2
		result := works(curr, weights)
		if result > D {
			// check upper half
			min = curr + 1
		} else {
			// check lower half
			max = curr - 1
		}
	}
}
