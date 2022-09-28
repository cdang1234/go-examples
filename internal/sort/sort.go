package sort

import (
	"fmt"
	"sort"
)

// implement custom comparator for sorting int32 slice. Golang std lib only has sorting for int types.
type byValueAsc []int32

func (f byValueAsc) Len() int {
	return len(f)
}

func (f byValueAsc) Less(i, j int) bool {
	return f[i] < f[j]
}

func (f byValueAsc) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

type byValueDesc []int32

func (f byValueDesc) Len() int {
	return len(f)
}

func (f byValueDesc) Less(i, j int) bool {
	return f[i] < f[j]
}

func (f byValueDesc) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

func Ascending() {
	// Write your code here
	arr := []int32{100, -1, 4, 9, 8, 600, 4, 3, 15, 75, 24, 8, -2}

	sort.Sort(byValueAsc(arr))

	fmt.Println(arr)
}

func Descending() {
	// Write your code here
	arr := []int32{100, -1, 4, 9, 8, 600, 4, 3, 15, 75, 24, 8, -2}

	sort.Sort(byValueDesc(arr))

	fmt.Println(arr)
}

// family := []struct {
// 	Name string
// 	Age  int
// }{
// 	{"Alice", 23},
// 	{"David", 2},
// 	{"Eve", 2},
// 	{"Bob", 25},
// }

// // Sort by age, keeping original order or equal elements.
// sort.SliceStable(family, func(i, j int) bool {
// 	return family[i].Age < family[j].Age
// })
// fmt.Println(family) // [{David 2} {Eve 2} {Alice 23} {Bob 25}]
