package leetcode

// 删除链表的倒数第N个结点
// https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/

// 快慢指针
func removeNthFromEndFastSlow(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head} // 哑节点，处理删除头节点的特殊情况
	fast, slow := head, dummy
	for i := 0; i < n; i++ { // fast先走n次
		fast = fast.Next
	}
	for ; fast != nil; fast = fast.Next { // 快慢指针同时移动
		slow = slow.Next
	}
	slow.Next = slow.Next.Next // 删除倒数第n个结点
	return dummy.Next
}

// 两次遍历：先计算链表长度再删除
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	length := 0
	cur := head
	for cur != nil {
		length++
		cur = cur.Next
	}

	dummy := &ListNode{0, head}
	cur = dummy
	for i := 0; i < length-n; i++ {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	return dummy.Next
}
