package leetcode

import "strings"

func longestCommonPrefix(strs []string) string {
	var a rune
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	for j := 0; j < len(strs[0]); j++ {
		a = []rune(strs[0])[j]
		for i := 1; i < len(strs); i++ {
			if len(strs[i]) <= j || a != []rune(strs[i])[j] {
				return strs[0][:j]
			}
		}
	}
	return strs[0]
}

func longestCommonPrefix2(strs []string) string {
	if len(strs) < 1 {
		return ""
	}
	prefix := strs[0]
	for _, v := range strs {
		for strings.Index(v, prefix) != 0 {
			if len(prefix) == 0 {
				return ""
			}
			prefix = prefix[:len(prefix)-1]
		}
	}
	return prefix
}
