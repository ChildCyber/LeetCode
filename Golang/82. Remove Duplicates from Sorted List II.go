package leetcode

// 删除排序链表中的重复元素 II
// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii/
// 一次遍历
func deleteDuplicates82(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	dummy := &ListNode{0, head}
	cur := dummy
	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val { // 当前 cur.next 与 cur.next.next 对应的元素相同，需要将 cur.next 以及所有后面拥有相同元素值的链表节点全部删除
			x := cur.Next.Val
			for cur.Next != nil && cur.Next.Val == x { // 删除连续出现的重复元素
				cur.Next = cur.Next.Next
			}
		} else { // 当前 cur.next 与 cur.next.next 对应的元素不相同，那么说明链表中只有一个元素值为 cur.next 的节点
			cur = cur.Next
		}
	}

	return dummy.Next
}

// 递归
func deleteDuplicates82Rec(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	if head.Val != head.Next.Val {
		head.Next = deleteDuplicates82Rec(head.Next)
	} else { // 出现重复元素
		move := head.Next // 标记需要移除的元素
		for move != nil && head.Val == move.Val {
			move = move.Next
		}
		return deleteDuplicates82Rec(move)
	}

	return head
}

// 计数+两次遍历
func deleteDuplicates82Two(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	dummy := &ListNode{0, head}

	cur := head
	counter := make(map[int]int)
	for cur != nil {
		counter[cur.Val]++
		cur = cur.Next
	}

	cur = dummy
	for cur != nil && cur.Next != nil {
		if counter[cur.Next.Val] != 1 {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}

	return dummy.Next
}
