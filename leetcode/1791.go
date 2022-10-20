package leetcode

func findCenter(edges [][]int) int {
	m, n := edges[0][0], edges[0][1]
	for _, v := range edges[1] {
		switch v {
		case m:
			return m
		case n:
			return n
		}
	}
	return 0
}
