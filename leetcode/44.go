package leetcode

/*
输入:
s = "aa"
p = "a"
输出: false
解释: "a" 无法匹配 "aa" 整个字符串。

'?' 可以匹配任何单个字符。
'*' 可以匹配任意字符串（包括空字符串）。
*/
func isMatch44(s string, p string) bool {
	if p == "*" {
		return true
	}
	var i, j int
	for i < len(s) {
		switch p[j] {
		case '*':
			p = p[:j] + "*" + p[j:]
		case '?':
			i++
			j++
			continue
		default:
		}
		if !equalsTo(s[i], p[j]) {
			return false
		}
		i++
		j++
	}
	if j < len(p) {
		return false
	}
	return true
}

func equalsTo(u uint8, u2 uint8) bool {
	switch u2 {
	case '*', '?':
		return true
	default:
		return u == u2
	}
}

func isMatch4(s string, p string) bool {
	m, n := len(s), len(p)
	// 我们用 dp[i][j] 表示字符串 ss 的前 ii 个字符和模式 pp 的前 jj 个字符是否能匹配。
	dp := make([][]bool, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	for i := 1; i <= n; i++ {
		if p[i-1] == '*' {
			dp[0][i] = true
		} else {
			break
		}
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				dp[i][j] = dp[i][j-1] || dp[i-1][j]
			} else if p[j-1] == '?' || s[i-1] == p[j-1] {
				dp[i][j] = dp[i-1][j-1]
			}
		}
	}
	return dp[m][n]
}
