package leetcode

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	a := len(nums1)
	b := len(nums2)
	if (a+b)%2 == 0 {
		mid1, mid2 := findEvenMedianInArrays(nums1, nums2)
		return calculateEvenMedian(mid1, mid2)
	} else {
		return float64(findOddMedianInArrays(nums1, nums2))
	}
}

func findOddMedianInArrays(nums1 []int, nums2 []int) int {
	a := len(nums1)
	b := len(nums2)
	m, n, t := 0, 0, 0
	res := 0
	for {
		if t > (a+b)/2 {
			break
		}
		if m >= a {
			res = nums2[n+(a+b)/2-t]
			break
		}
		if n >= b {
			res = nums1[m+(a+b)/2-t]
			break
		}
		if nums1[m] < nums2[n] {
			res = nums1[m]
			m++
		} else {
			res = nums2[n]
			n++
		}
		t = m + n
	}
	return res
}

func findEvenMedianInArrays(nums1 []int, nums2 []int) (int, int) {
	a := len(nums1)
	b := len(nums2)
	m, n, t := 0, 0, 0
	res1, res2 := 0, 0
	for {
		if t > (a+b)/2 {
			break
		}

		if m >= a {
			if n+(a+b)/2-t-1 >= 0 {
				res1 = nums2[n+(a+b)/2-t-1]
			} else {
				res1 = res2
			}
			res2 = nums2[n+(a+b)/2-t]
			break
		}
		if n >= b {
			if m+(a+b)/2-t-1 >= 0 {
				res1 = nums1[m+(a+b)/2-t-1]
			} else {
				res1 = res2
			}
			res2 = nums2[m+(a+b)/2-t]
			break
		}
		res1 = res2
		if nums1[m] < nums2[n] {
			res2 = nums1[m]
			m++
		} else {
			res2 = nums2[n]
			n++
		}
		t = m + n
	}
	return res1, res2
}

func calculateEvenMedian(a, b int) float64 {
	return float64(a+b) / 2
}
