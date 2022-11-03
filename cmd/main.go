package main

import "fmt"

func main() {
	pointers := []int{0, 1, 2, 3}
	k := 2
	fmt.Println(genPointers(pointers, k))
}

func genPointers(nums []int, k int) [][]int {
	result := [][]int{}

	if k == 1 {
		for i := 0; i < len(nums); i++ {
			result = append(result, []int{i})
		}
		return result
	}

	for i := 0; i < len(nums)-2; i++ {
		nums2 := make([]int, len(nums))

		// operate on copy of nums
		copy(nums2, nums)
		n := nums2[i+1:]
		fmt.Println(n)
		fmt.Println(i)

		pieces := genPointers(n, k-1)
		for _, piece := range pieces {
			result = append(result, append([]int{i}, piece...))
		}
	}

	return result
}
