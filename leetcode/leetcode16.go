package leetcode

import "sort"

const (
	IntMax = int(^uint(0) >> 1)
)

func threeSumClosest(nums []int, target int) int {
	// {2,3,4,6,7,12}, target=12
	sort.SliceStable(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	min, total, rs := IntMax, IntMax, IntMax
	for i := 0; i < len(nums); i++ {
		pilot := nums[i]
		left, right := i+1, len(nums)-1
		for {
			if left >= right {
				break
			}
			total = pilot + nums[left] + nums[right]
			if target < total {
				if min > total-target {
					rs = total
					min = total - target
				}
				right--
				continue
			}
			if min > target-total {
				rs = total
				min = target - total
			}
			left++
		}
	}
	return rs
}
