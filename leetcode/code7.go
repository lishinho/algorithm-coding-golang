package leetcode

func reverse(x int) int {
	var res int
	for {
		res = res*10 + x%10
		x = x / 10
		if res < -2147483648 || res >= 2147483648 {
			return 0
		}
		if x == 0 {
			break
		}
	}
	return res
}
