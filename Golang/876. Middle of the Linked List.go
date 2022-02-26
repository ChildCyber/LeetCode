package leetcode

// 链表的中间结点
// https://leetcode-cn.com/problems/middle-of-the-linked-list/
// 快慢指针
func middleNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	p1, p2 := head, head
	for p2.Next != nil && p2.Next.Next != nil {
		p1 = p1.Next      // 慢指针
		p2 = p2.Next.Next // 快指针
	}

	length := 0
	cur := head
	for cur != nil {
		length++
		cur = cur.Next
	}

	if length%2 == 0 { // 链表长度为偶数
		return p1.Next
	}
	return p1
}
