package leetcode

import "math"

/*
nums = [2,3,1,1,4]
输出: 2
解释: 跳到最后一个位置的最小跳跃数是 2。
从下标为 0 跳到下标为 1 的位置，跳1步，然后跳3步到达数组的最后一个位置。
*/

func jump(nums []int) int {
	// dp[i] 在第i格上能跳到最后的最少跳跃数
	// dp[i] = min{dp[i+m]+1}, m=nums[i]
	// return dp[0]
	if len(nums) <= 1 {
		return 0
	}
	n := len(nums)
	dp := make([]int, n)
	dp[n-1] = 0
	var f func(i, m int) int
	f = func(i, m int) int {
		res := math.MaxInt - 1
		for j := 1; j <= m; j++ {
			if i+j > n-1 {
				break
			}
			res = min(res, dp[i+j]+1)
		}
		return res
	}
	for i := n - 2; i >= 0; i-- {
		dp[i] = f(i, nums[i])
	}
	return dp[0]
}

func jump45(nums []int) int {
	position := len(nums) - 1
	steps := 0
	for position > 0 {
		for i := 0; i < position; i++ {
			if i+nums[i] >= position {
				position = i
				steps++
				break
			}
		}
	}
	return steps
}
