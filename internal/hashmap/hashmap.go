package hashmap

import (
	"fmt"
	"sort"
)

func GetKeys() {
	mymap := make(map[int]string)

	for k, v := range mymap {
		fmt.Printf("key[%s] value[%s]\n", k, v)
	}
}

func GetSortedKeys() {
	mapper := make(map[int]string)
	mapper[0] = "a"
	mapper[1] = "b"
	mapper[2] = "c"
	for key := range mapper {
		fmt.Printf("%v \n", key)
	}

	i := 0
	keys := make([]int, len(mapper))
	for k := range mapper {
		keys[i] = k
		i++
	}

	sort.Ints(keys)
	fmt.Println(keys)
}
