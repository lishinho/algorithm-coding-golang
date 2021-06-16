package main

import (
	"fmt"
	"sort"
)

func sortArray() {
	tmp := []int{6, 2, 3, 4, 3, 2, 5, 43}
	// 从大到小
	//sort.SliceStable(tmp, func(i, j int) bool {
	//	return tmp[i] > tmp[j]
	//})

	// 从小到大
	sort.Ints(tmp)
	fmt.Println(tmp)
}
