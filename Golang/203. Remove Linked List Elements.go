package leetcode

// 移除链表元素
// https://leetcode.cn/problems/remove-linked-list-elements/
// 迭代
func removeElements(head *ListNode, val int) *ListNode {
	// 由于链表的头节点 head 有可能会被删除，因此创建哑节点 dummyHead
	dummyHead := &ListNode{Next: head}
	for tmp := dummyHead; tmp.Next != nil; {
		if tmp.Next.Val == val {
			tmp.Next = tmp.Next.Next
		} else {
			tmp = tmp.Next
		}
	}
	return dummyHead.Next
}

// 递归
func removeElementsRec(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}

	head.Next = removeElementsRec(head.Next, val)
	if head.Val == val {
		return head.Next
	}

	return head
}
