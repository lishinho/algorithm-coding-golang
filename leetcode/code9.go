package leetcode

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	var arr []int8
	for {
		if x == 0 {
			break
		}
		arr = append(arr, int8(x%10))
		x = x / 10
	}
	i, j := 0, len(arr)-1
	for {
		if i >= j {
			return true
		}
		if arr[i] != arr[j] {
			return false
		}
		i++
		j--
	}
}
