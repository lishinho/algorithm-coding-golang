package main

import (
	"fmt"
)

func main() {
	fmt.Println(AvToBv(602109295))
	fmt.Println(BvToAv("BV1bg411A7pa"))
	//前闭后开()
	//fmt.Println(strconv.FormatInt(100, 2))
	grpcJsonGet()
}
