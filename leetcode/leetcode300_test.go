package leetcode

import (
	"testing"
	"time"
)

func TestLeetcode300(t *testing.T) {
	cur := time.Now()
	test1 := []int{10, 9, 2, 5, 3, 7, 101, 18}
	if lengthOfLTS(test1) != 4 {
		t.Fail()
	}
	var test2 []int
	if lengthOfLTS(test2) != 0 {
		t.Fail()
	}
	test3 := []int{2, 9, 73, -5, 0}
	if lengthOfLTS(test3) != 3 {
		t.Fail()
	}
	t.Logf("success in %v", time.Since(cur))
}
