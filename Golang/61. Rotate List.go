package leetcode

// 旋转链表
// https://leetcode-cn.com/problems/rotate-list/
// 闭合为环
func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil || head.Next == nil {
		return head
	}

	n := 1
	iter := head
	for iter.Next != nil {
		iter = iter.Next
		n++
	}
	add := n - k%n
	// 当链表长度不大于 1，或者 k 为 n 的倍数时，新链表将与原链表相同
	if add == n {
		return head
	}

	iter.Next = head
	for add > 0 {
		iter = iter.Next
		add--
	}
	ret := iter.Next
	iter.Next = nil
	return ret
}
