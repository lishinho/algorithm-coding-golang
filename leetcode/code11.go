package leetcode

func maxArea(height []int) int {
	if len(height) == 0 {
		return 0
	}
	var res int
	for i, j := 0, len(height)-1; i < j; {
		if height[i] < height[j] {
			res = max(res, height[i]*(j-i))
			i++
			continue
		}
		res = max(res, height[j]*(j-i))
		j--
	}
	return res
}
