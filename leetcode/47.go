package leetcode

import (
	"sort"
)

/*
输入：nums = [1,1,2]
输出：
[[1,1,2],
 [1,2,1],
 [2,1,1]]
*/

func permuteUnique2(nums []int) (res [][]int) {
	n := len(nums)
	var backtrace func(start int, is bool)
	backtrace = func(start int, is bool) {
		if start == n-1 && !is {
			res = append(res, append([]int(nil), nums...))
			return
		}
		if nums[start] == nums[start+1] {
			backtrace(start+1, true)
		}
		for i := start; i < n; i++ {
			nums[start], nums[i] = nums[i], nums[start]
			backtrace(start+1, false)
			nums[start], nums[i] = nums[i], nums[start]
		}
	}
	backtrace(0, false)
	return res
}

func permuteUnique(nums []int) (ans [][]int) {
	sort.Ints(nums)
	n := len(nums)
	var perm []int
	vis := make([]bool, n)
	var backtrack func(int)
	backtrack = func(idx int) {
		if idx == n {
			ans = append(ans, append([]int(nil), perm...))
			return
		}
		for i, v := range nums {
			if vis[i] || i > 0 && !vis[i-1] && v == nums[i-1] {
				continue
			}
			perm = append(perm, v)
			vis[i] = true
			backtrack(idx + 1)
			vis[i] = false
			perm = perm[:len(perm)-1]
		}
	}
	backtrack(0)
	return
}
