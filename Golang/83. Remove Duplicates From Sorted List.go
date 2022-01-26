package leetcode

// 删除排序列表中的重复元素
// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/
// 一次遍历
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}

	cur := head
	for cur.Next != nil { // 只需要遍历到链表的最后一个节点，而不需要遍历完整个链表
		if cur.Next.Val == cur.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}
