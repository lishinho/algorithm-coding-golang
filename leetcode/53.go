package leetcode

func maxSubArray(nums []int) int {
	// dp[i] = 从i开始最大和的连续子数组
	// dp[i] = max{dp[i-1]+nums[i], dp[i-1]}
	n := len(nums)
	dp := make([]int, n)
	dp[n-1] = nums[n-1]

	res := dp[n-1]
	for i := n - 2; i >= 0; i-- {
		dp[i] = max(dp[i+1]+nums[i], nums[i])
		res = max(res, dp[i])
	}
	return res
}
