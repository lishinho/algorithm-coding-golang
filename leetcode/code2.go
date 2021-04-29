package leetcode

/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l0 := &ListNode{}
	cur := l0
	incre := 0
	for {
		if l1 == nil || l2 == nil {
			break
		}
		cur.Next, incre = calculateNextListNode(l1.Val, l2.Val, incre)
		cur = cur.Next
		l1 = l1.Next
		l2 = l2.Next
	}
	for {
		if l1 == nil {
			break
		}
		cur.Next, incre = calculateNextListNode(l1.Val, 0, incre)
		cur = cur.Next
		l1 = l1.Next
	}
	for {
		if l2 == nil {
			break
		}
		cur.Next, incre = calculateNextListNode(0, l2.Val, incre)
		cur = cur.Next
		l2 = l2.Next
	}
	if incre != 0 {
		cur.Next, incre = calculateNextListNode(0, 0, incre)
	}
	return l0.Next
}

func calculateNextListNode(v1, v2, incre int) (*ListNode, int) {
	nextVal := v1 + v2 + incre
	if nextVal > 9 {
		return &ListNode{Val: nextVal % 10}, 1
	}
	return &ListNode{Val: nextVal}, 0
}
