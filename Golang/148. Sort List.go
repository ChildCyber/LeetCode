package leetcode

// 排序链表
// https://leetcode-cn.com/problems/sort-list/

// 合并两个有序链表
func merge148(head1, head2 *ListNode) *ListNode { // 合并两个排序后的子链表
	dummy := &ListNode{}
	cur, temp1, temp2 := dummy, head1, head2
	for temp1 != nil && temp2 != nil {
		if temp1.Val <= temp2.Val {
			cur.Next = temp1
			temp1 = temp1.Next
		} else {
			cur.Next = temp2
			temp2 = temp2.Next
		}
		cur = cur.Next
	}

	if temp1 != nil {
		cur.Next = temp1
	} else if temp2 != nil {
		cur.Next = temp2
	}

	return dummy.Next
}

// 使用快慢指针找到链表中点
func findMiddle(head *ListNode) *ListNode {
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// 递归（自顶向下归并排序）
// 时间复杂度：O(n log n)
// 空间复杂度：O(log n)
func sortListTopDown(head *ListNode) *ListNode {
	// 基线条件：空链表或单节点链表已经有序
	if head == nil || head.Next == nil {
		return head
	}

	// 找到链表中点
	mid := findMiddle(head)

	// 分割链表
	rightHead := mid.Next
	mid.Next = nil // 切断链表

	// 递归排序左右两部分
	left := sortListTopDown(head)
	right := sortListTopDown(rightHead)

	// 合并两个有序链表
	return merge148(left, right)
}

// 迭代（自底向上归并排序）
// 时间复杂度：O(n log n)
// 空间复杂度：O(1)
func sortListBottomUp(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 获取链表长度
	length := 0
	for node := head; node != nil; node = node.Next {
		length++
	}

	dummy := &ListNode{Next: head}
	// 子链表大小从1开始，每次翻倍
	for subLen := 1; subLen < length; subLen *= 2 {
		prev, curr := dummy, dummy.Next
		for curr != nil {
			// 获取第一个子链表（长度为subLen）
			left := curr
			for i := 1; i < subLen && curr.Next != nil; i++ {
				curr = curr.Next
			}

			// 获取第二个子链表（长度为subLen）
			right := curr.Next
			curr.Next = nil // 切断第一个子链表
			curr = right
			for i := 1; i < subLen && curr != nil && curr.Next != nil; i++ {
				curr = curr.Next
			}

			// 保存剩余部分
			var next *ListNode
			if curr != nil {
				next = curr.Next
				curr.Next = nil // 切断第二个子链表
			}

			// 合并两个子链表
			merged := merge148(left, right)
			prev.Next = merged
			// 移动到合并后链表的末尾
			for prev.Next != nil {
				prev = prev.Next
			}

			// 推进到下一组子链表
			curr = next
		}
	}

	return dummy.Next
}
