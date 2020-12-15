package leetcode

func lengthOfLTS(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}
	res := 1
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		res = max(res, dp[i])
	}
	return res

}
