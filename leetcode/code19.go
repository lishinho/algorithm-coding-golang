package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	l0, l1 := dummy, head
	for i := 0; i < n; i++ {
		if l1 == nil {
			return head
		}
		l1 = l1.Next
	}
	for {
		if l1 == nil {
			cur := l0.Next
			l0.Next = cur.Next
			cur.Next = nil
			break
		}
		l0 = l0.Next
		l1 = l1.Next
	}
	return dummy.Next
}
