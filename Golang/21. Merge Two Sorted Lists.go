package leetcode

// 合并两个有序链表
// https://leetcode-cn.com/problems/merge-two-sorted-lists/

// 递归
// 时间复杂度：O(m+n)
// 空间复杂度：O(m+n)
func mergeTwoListsRecursive(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		l1.Next = mergeTwoListsRecursive(l1.Next, l2)
		return l1
	}
	l2.Next = mergeTwoListsRecursive(l1, l2.Next)
	return l2
}

// 迭代
// 时间复杂度：O(m+n)
// 空间复杂度：O(1)
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{} // 哑节点，比较容易地返回合并后的链表
	head := dummy

	for list1 != nil && list2 != nil { // 都非空
		// 比较大小
		if list1.Val < list2.Val {
			head.Next = list1
			list1 = list1.Next
		} else {
			head.Next = list2
			list2 = list2.Next
		}
		head = head.Next
	}

	// 处理剩余的链表部分
	if list1 != nil {
		head.Next = list1
	} else {
		head.Next = list2
	}

	return dummy.Next
}
