package leetcode

var (
	intToRomanLimit = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	intToRomanToken = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
)

func intToRoman(num int) string {
	var res string
	for {
		if num == 0 {
			break
		}
		for i := range intToRomanLimit {
			calNum, calRes, ac := simplifyIntToRomanMeta(num, intToRomanLimit[i], res, intToRomanToken[i])
			num = calNum
			res = calRes
			if ac {
				break
			}
		}
	}
	return res
}

func simplifyIntToRomanMeta(num, i int, res, token string) (int, string, bool) {
	if num >= i {
		num -= i
		res = res + token
		return num, res, true
	}
	return num, res, false
}

var valueSymbols = []struct {
	value  int
	symbol string
}{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func intToRoman2(num int) string {
	roman := []byte{}
	for _, vs := range valueSymbols {
		for num >= vs.value {
			num -= vs.value
			roman = append(roman, vs.symbol...)
		}
		if num == 0 {
			break
		}
	}
	return string(roman)
}
