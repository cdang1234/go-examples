package search

import (
	"fmt"
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
