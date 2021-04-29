package leetcode

func longestPalindrome(s string) string {
	var res string
	for i := 0; i < len(s); i++ {
		res = maxLengthString(res, fixedDoubleCursorLongest(s, i))
		res = maxLengthString(fixedSingleCursorLongest(s, i), res)
	}
	return res
}

func maxLengthString(s1, s2 string) string {
	if len(s1) > len(s2) {
		return s1
	}
	return s2
}

func fixedSingleCursorLongest(s string, cur int) string {
	r := []rune(s)
	i, j := cur, cur
	for ; i >= 0 && j < len(r); i, j = i-1, j+1 {
		if r[i] != r[j] {
			break
		}
	}
	if i+1 >= 0 && j <= len(r) {
		return s[i+1 : j]
	}
	return ""
}

func fixedDoubleCursorLongest(s string, cur int) string {
	r := []rune(s)
	i, j := cur, cur+1
	for ; i >= 0 && j < len(r); i, j = i-1, j+1 {
		if r[i] != r[j] {
			break
		}
	}
	if i+1 >= 0 && j <= len(r) && i != cur {
		return s[i+1 : j]
	}
	return ""
}
