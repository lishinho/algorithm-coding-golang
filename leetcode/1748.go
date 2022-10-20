package leetcode

type obj struct {
	unMinus bool
}

func sumOfUnique(nums []int) int {
	m := make(map[int]*obj, len(nums))
	var res int
	for _, num := range nums {
		if v, ok := m[num]; ok && !v.unMinus {
			res -= num
			v.unMinus = true
		} else if !ok {
			res += num
			m[num] = &obj{false}
		}
	}
	return res
}
