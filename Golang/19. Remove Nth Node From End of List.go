package leetcode

// 删除链表的倒数第N个结点
// https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/
// 快慢指针
func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	// https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/solution/dong-hua-tu-jie-leetcode-di-19-hao-wen-ti-shan-chu/
	dummy := &ListNode{Next: head}
	first, second := head, dummy
	for i := 0; i < n; i++ { // 先走n次
		first = first.Next
	}
	for ; first != nil; first = first.Next {
		second = second.Next
	}
	second.Next = second.Next.Next // 删除倒数第N个结点
	return dummy.Next
}

// 计算链表长度
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	length := 0
	cur := head
	for ; head != nil; head = head.Next {
		length++
	}

	dummy := &ListNode{0, head}
	cur = dummy
	for i := 0; i < length-n; i++ {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	return dummy.Next
}
