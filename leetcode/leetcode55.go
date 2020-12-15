package leetcode

func canJump(nums []int) bool {
	n := len(nums)
	rightMost := 0
	for i := 0; i < n-1; i++ {
		if i > rightMost {
			return false
		}
		rightMost = max(rightMost, i+nums[i])
		if rightMost >= n-1 {
			return true
		}
	}
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
