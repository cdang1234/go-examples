package slice

import "fmt"

func Traverse2DSlice() {
	arr := [][]int{[]int{1, 2, 3, 4}, []int{5, 6, 7, 8}, []int{9, 10, 11, 12}, []int{13, 14, 15, 16}}

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Println(arr[i][j])
		}
	}

	// output:
	// 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16
	// i accesses outer array, j accesses inner array
}

func Reverse() {
	a := []int{1, 2, 3, 4}
	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
}

func SliceMechanics() {
	// paritioning
	s := []int{1, 2, 3}
	fmt.Println(s[:1]) // [1]
	fmt.Println(s[2:]) // [3]

	fmt.Println(s[:2]) // [1, 2]
	fmt.Println(s[3:]) // []

	// underlying data model
	z := s[2:] // z points to 3rd index of s
	z[0] = 4   // slices are pointers to underlying values

	fmt.Println(s) // 1 2 4
}
