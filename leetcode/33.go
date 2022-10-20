package leetcode

import "sort"

//输入：nums = [4,5,6,7,0,1,2], target = 0
//输出：4

func search33(nums []int, target int) int {
	var num []int
	copy(num, nums)
	sort.Ints(nums)
	var res int
	var f func(nums []int)

	f = func(nums []int) {
		i, j := 0, len(nums)
		for i < j {
			s := int(uint(i+j) >> 1)
			if nums[s] == target {
				res = s
				return
			} else if nums[s] > target {
				j = s
			} else {
				i = s + 1
			}
		}

	}

	f(nums)
	target = res
	f(num)
	return res
}
