package leetcode

func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}
	for i := len(nums) - 1; i > 1; i-- {
		if nums[i] > nums[i-1] {
			nums[i], nums[i-1] = nums[i-1], nums[i]
			return
		}
	}
	if nums[0] < nums[1] {
		nums[0], nums[1] = nums[1], nums[0]
		return
	}
	nums[0], nums[len(nums)-1] = nums[len(nums)-1], nums[0]
}
