package leetcode

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{}
	dummy.Next = head
	prev := dummy
	left := head
	right := head.Next
	for {
		if right == nil {
			break
		}
		swap(prev, left, right)
		prev = left
		left = left.Next
		if left == nil {
			break
		}
		right = left.Next
	}
	return dummy.Next
}

func swap(prev, left, right *ListNode) {
	prev.Next = right
	left.Next = right.Next
	right.Next = left
}
