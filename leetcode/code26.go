package leetcode

// 无序数组仍然受用
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	m := map[int]struct{}{}
	for i := 0; i < len(nums); i++ {
		if _, ok := m[nums[i]]; !ok {
			m[nums[i]] = struct{}{}
		} else {
			nums = append(nums[:i-1], nums[i:]...)
			i--
		}
	}
	return len(nums)
}
