package leetcode

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	dp := make([]int, len(s))

	length := len(s)
	dp[length-1] = 1
	for i := length - 2; i >= 0; i-- {
		dp[i] = maxNumber(dp[i+1], fromNewSubstring(s, i))
	}
	return dp[0]
}

func maxNumber(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func fromNewSubstring(s string, i int) int {
	m := map[uint8]bool{}
	for j := i; j < len(s); j++ {
		if _, ok := m[s[j]]; ok {
			return len(s[i:j])
		}
		m[s[j]] = true
	}
	return len(s[i:])
}
