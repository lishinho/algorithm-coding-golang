package leetcode

func mergeKLists(lists []*ListNode) *ListNode {
	// 两个两个合并
	if len(lists) < 1 {
		return nil
	}
	res := lists[0]
	if len(lists) == 1 {
		return res
	}
	for i := 1; i < len(lists); i++ {
		res = mergeTwoLists(res, lists[i])
	}
	return res
}
