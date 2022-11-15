package main

import (
	"encoding/json"
	"fmt"
)

func arrayToJson() {
	arr := []int64{1, 2, 3, 4, 5}
	m := map[string][]int64{
		"str": arr,
		"abc": arr,
	}
	bt, _ := json.Marshal(m)
	fmt.Printf("%s", bt[1:len(bt)-1])
}

func arrayWithk() {
	arr := []int64{1, 2, 3, 4, 5}
	k := 3
	fmt.Printf("%+v", arr[:k])
	fmt.Printf("%+v", arr[k:])
}
