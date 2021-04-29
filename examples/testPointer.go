package main

import "fmt"

type student struct {
	name  string
	score int32
}

//func main() {
//	stu1 := student{
//		name:  "lishinho",
//		score: 50,
//	}
//	// attention
//	modifyScore(&stu1, 10)
//	fmt.Printf("stu1 data: %v", stu1)
//}

func modifyScore(stu *student, score int32) {
	*stu = student{
		name:  "hh",
		score: score,
	}
	fmt.Printf("stu in is %v || %v", *stu, stu)
}
