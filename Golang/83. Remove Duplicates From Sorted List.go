package leetcode

// 删除排序列表中的重复元素
// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/
// 一次遍历
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	cur := head
	// 遍历到链表的最后一个节点时，cur.next为空节点
	for cur.Next != nil { // 只需要遍历到链表的最后一个节点，而不需要遍历完整个链表
		if cur.Next.Val == cur.Val { // cur与cur.next对应的元素相同，将cur.next从链表中移除，删掉下一个重复元素
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}

// 递归
func deleteDuplicatesRec(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	head.Next = deleteDuplicatesRec(head.Next)
	if head.Val == head.Next.Val {
		return head.Next
	}
	return head
}
