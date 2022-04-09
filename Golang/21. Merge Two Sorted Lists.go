package leetcode

// 合并两个有序链表
// https://leetcode-cn.com/problems/merge-two-sorted-lists/
// 递归
func mergeTwoListsRec(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		l1.Next = mergeTwoListsRec(l1.Next, l2)
		return l1
	}
	l2.Next = mergeTwoListsRec(l1, l2.Next)
	return l2
}

// 迭代
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{} // 额外的头节点，比较容易地返回合并后的链表
	cur := head         // 变换cur节点的Next指针
	for l1 != nil || l2 != nil {
		if l1 != nil && l2 != nil { // 都非空
			if l1.Val < l2.Val { // 比较大小
				cur.Next = l1
				l1 = l1.Next
			} else {
				cur.Next = l2
				l2 = l2.Next
			}
			cur = cur.Next // cur指针向后走
		} else if l1 != nil { // l1非空，l2空
			cur.Next = l1
			break
		} else { // l2非空，l1空
			cur.Next = l2
			break
		}
	}
	return head.Next
}
