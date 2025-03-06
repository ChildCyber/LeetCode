package leetcode

// 链表的中间结点
// https://leetcode-cn.com/problems/middle-of-the-linked-list/

// 快慢指针
func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}
