package leetcode

import "math"

func myAtoi(s string) int {
	max := int(math.Pow(2, 31))
	n := 0
	for i, ch := range []byte(s) {
		ch -= '0'
		if ch > 0 && ch <= 9 {
			n = postMyAtoi(s[i:])
			if n > max {
				n = max
			}
			if i != 0 && s[i-1] == '-' {
				n = -n
			}
			break
		}
	}
	return n
}

func postMyAtoi(s string) int {
	n := 0
	for _, ch := range []byte(s) {
		ch -= '0'
		if ch < 0 || ch > 9 {
			break
		}
		n = n*10 + int(ch)
	}
	return n
}
