package main

import (
	"fmt"
	"sort"
)

func sortArray() {
	tmp := []int{6, 2, 3, 4, 3, 2, 5, 43}
	sort.SliceStable(tmp, func(i, j int) bool {
		return tmp[i] > tmp[j]
	})
	fmt.Println(tmp)
}
