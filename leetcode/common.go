package leetcode

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type LRUCache2 struct {
	capcity int
}

func (*LRUCache) name() {

}
