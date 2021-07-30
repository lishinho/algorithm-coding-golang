package leetcode

var brackets = map[string]string{
	")": "(",
	"]": "[",
	"}": "{",
}

func isValid(s string) bool {
	var stack []string
	for i := 0; i < len(s); i++ {
		val, ok := brackets[string(s[i])]
		if !ok {
			stack = append(stack, string(s[i]))
			continue
		}
		if len(stack) == 0 || stack[len(stack)-1] != val {
			return false
		}
		stack = stack[:len(stack)-1]
	}
	if len(stack) == 0 {
		return true
	}
	return false
}
