package leetcode

import "sort"

func threeSum(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	var res [][]int
	pilot := -1000001
	// sort nums[i], nums[j], pilot-nums[j]
	sort.SliceStable(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	for i := 0; i < len(nums); i++ {
		if pilot == -nums[i] {
			continue
		}
		pilot = -nums[i]
		sMap := make(map[int]bool)
		for j := i + 1; j < len(nums); j++ {
			// map[pilot-nums[j]] = bool
			if val, ok := sMap[nums[j]]; ok {
				if !val {
					// 存在非第一次
					continue
				}
				// 存在第一次
				res = append(res, []int{nums[i], pilot - nums[j], nums[j]})
				//sMap[pilot-nums[j]] = false
				sMap[nums[j]] = false
				continue
			}
			// 不存在
			sMap[pilot-nums[j]] = true
		}
	}
	return res
}

func threeSum2(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		if nums[i] <= nums[j] {
			return true
		}
		return false
	})
	res := make([][]int, 0)
	n := len(nums)
	for i := 0; i < n && nums[i] <= 0; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		t := -nums[i]
		l, r := i+1, n-1
		for l < r {
			sum := nums[l] + nums[r]
			if sum < t {
				l++
			} else if sum > t {
				r--
			} else {
				res = append(res, []int{-t, nums[l], nums[r]})
				for l < r && nums[l+1] == nums[l] {
					l++
				}
				for l < r && nums[r-1] == nums[r] {
					r--
				}
				l++
				r--
			}
		}
	}
	return res

}
