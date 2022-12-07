package leetcode

func partition(s string) (ans [][]string) {
	n := len(s)
	f := make([][]bool, n)
	for i := range f {
		f[i] = make([]bool, n)
		for j := range f[i] {
			f[i][j] = true
		}
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			f[i][j] = s[i] == s[j] && f[i+1][j-1]
		}
	}
	var tmp []string
	var backtrace func(cur int)
	backtrace = func(cur int) {
		if cur >= n {
			ans = append(ans, append([]string(nil), tmp...))
			return
		}
		for i := cur; i < n; i++ {
			if f[cur][i] {
				tmp = append(tmp, s[cur:i+1])
				backtrace(i + 1)
				tmp = tmp[:len(tmp)-1]
			}
		}
	}
	backtrace(0)
	return ans
}
