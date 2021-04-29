package leetcode

func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for i, v := range nums {
		if storedId, ok := m[v]; ok {
			return []int{storedId, i}
		}
		m[target-v] = i
	}
	return []int{}
}
