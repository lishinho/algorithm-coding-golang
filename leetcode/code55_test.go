package leetcode

import "testing"

func TestLeetcode55(t *testing.T) {
	test1 := []int{1, 2, 3, 4, 5}
	if !canJump(test1) {
		t.Fail()
	}
	t.Logf("success with %v", test1)
}
