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
