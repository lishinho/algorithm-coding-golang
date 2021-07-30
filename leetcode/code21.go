package leetcode

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for {
		if l1 == nil || l2 == nil {
			break
		}
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
			cur = cur.Next
			continue
		}
		cur.Next = l2
		l2 = l2.Next
		cur = cur.Next
	}
	for {
		if l1 == nil {
			break
		}
		cur.Next = l1
		l1 = l1.Next
		cur = cur.Next
	}
	for {
		if l2 == nil {
			break
		}
		cur.Next = l2
		l2 = l2.Next
		cur = cur.Next
	}
	return dummy.Next
}
