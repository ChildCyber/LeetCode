package leetcode

// 分隔链表
// https://leetcode.cn/problems/partition-list/

// 双链表
func partition(head *ListNode, x int) *ListNode {
	// small链表按顺序存储所有小于x的节点
	small := &ListNode{}
	smallHead := small
	// large链表按顺序存储所有大于等于x的节点
	large := &ListNode{}
	largeHead := large

	// 历完原链表后，将small链表尾节点指向large链表的头节点
	for head != nil {
		if head.Val < x {
			small.Next = head
			small = small.Next
		} else {
			large.Next = head
			large = large.Next
		}
		head = head.Next
	}
	large.Next = nil
	small.Next = largeHead.Next

	return smallHead.Next
}
