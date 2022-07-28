package leetcode

/*
编写一个高效的算法来搜索mxn矩阵 matrix 中的一个目标值 target 。该矩阵具有以下特性：

每行的元素从左到右升序排列。
每列的元素从上到下升序排列。
*/
//[]int

func searchMatrix(matrix [][]int, target int) bool {
	for _, row := range matrix {
		t := searchInts(row, target)
		if t < len(row) && row[t] == target {
			return true
		}
	}
	return false
}

func searchInts(a []int, x int) int {
	return search(len(a), func(i int) bool { return a[i] >= x })
}

func search(n int, f func(int) bool) int {
	i, j := 0, n
	for i < j {
		h := int(uint(i+j) >> 1)
		if !f(h) {
			i = h + 1
		} else {
			j = h
		}
	}
	return i
}
