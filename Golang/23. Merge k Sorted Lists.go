package leetcode

// 合并K个升序链表
// https://leetcode-cn.com/problems/merge-k-sorted-lists/
// 分治 + 递归
// 借助分治的思想，把 K 个有序链表两两合并即可
func mergeKLists(lists []*ListNode) *ListNode {
	length := len(lists)
	if length < 1 {
		return nil
	}
	if length == 1 {
		return lists[0]
	}
	// 分治
	num := length / 2
	left := mergeKLists(lists[:num])
	right := mergeKLists(lists[num:])
	// 合并两条升序链表
	return mergeTwoLists23(left, right)
}

func mergeTwoLists23(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists23(l1.Next, l2)
		return l1
	}
	l2.Next = mergeTwoLists23(l1, l2.Next)
	return l2
}
