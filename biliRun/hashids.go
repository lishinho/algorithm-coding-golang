package main

import (
	"fmt"
	"github.com/speps/go-hashids"
)

// 测试hashids的可用性

func hash() {
	hd := hashids.NewData()
	hd.Salt = "this is my salt"
	hd.MinLength = 12
	h, _ := hashids.NewWithData(hd)
	e, _ := h.Encode([]int{45})
	fmt.Println(e)

	d, _ := h.DecodeWithError(e)
	fmt.Println(d)
}
