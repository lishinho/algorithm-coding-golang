package leetcode

func convert(s string, numRows int) string {
	if numRows == 0 {
		return ""
	}
	if numRows == 1 {
		return s
	}
	r := []rune(s)
	var res []rune
	pilot := numRows*2 - 2
	for i := 0; i < numRows; i++ {
		for j := 0; ; j++ {
			if pilot*j-i >= len(r) {
				break
			}
			if pilot*j-i >= 0 {
				res = append(res, r[pilot*j-i])
			}
			if pilot*j+i >= len(r) {
				break
			}
			if i > 0 && i < numRows-1 && pilot*j+i >= 0 {
				res = append(res, r[pilot*j+i])
			}
		}
	}
	return string(res)
}
