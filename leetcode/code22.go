package leetcode

func generateParenthesis(n int) []string {
	strs := make([]string, 0)
	generate(&strs, "", 0, 0, n)
	return strs
}

func generate(strs *[]string, str string, leftCnt int, rightCnt int, n int) {
	if leftCnt > n || rightCnt > n {
		return
	}
	if leftCnt == n && rightCnt == n {
		*strs = append(*strs, str)
	}
	if leftCnt >= rightCnt {
		generate(strs, str+"(", leftCnt+1, rightCnt, n)
		generate(strs, str+")", leftCnt, rightCnt+1, n)
	}
}
