package leetcode

// 删除链表的倒数第N个结点
// https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/
// 快慢指针
func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	// https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/solution/dong-hua-tu-jie-leetcode-di-19-hao-wen-ti-shan-chu/
	dummyHead := &ListNode{Next: head}
	preSlow, slow, fast := dummyHead, head, head
	for fast != nil {
		if n <= 0 { // fast & slow一起走
			preSlow = slow
			slow = slow.Next
		}
		// fast先走n次
		n--
		fast = fast.Next
	}
	preSlow.Next = slow.Next // 删除倒数第N个结点
	return dummyHead.Next
}
