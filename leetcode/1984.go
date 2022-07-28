package leetcode

import (
	"sort"
)

func minimumDifference(nums []int, k int) int {
	sort.Ints(nums)
	ans := 1<<31 - 1
	for i, num := range nums[:len(nums)-k+1] {
		ans = min(ans, nums[i+k-1]-num)
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
