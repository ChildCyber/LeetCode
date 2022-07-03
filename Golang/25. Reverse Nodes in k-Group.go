package leetcode

// K 个一组翻转链表
// https://leetcode-cn.com/problems/reverse-nodes-in-k-group/
//func reverseGroup(head *ListNode, k int) *ListNode {
//	node := head
//	for i := 0; i < k; i++ {
//		if node == nil {
//			return head
//		}
//		node = node.Next
//	}
//	newHead := reverse(head, node)
//	head.Next = reverseGroup(node, k)
//	return newHead
//}
//
//func reverse(first *ListNode, last *ListNode) *ListNode { // 反转两个链表节点
//	prev := last
//	for first != last {
//		tmp := first.Next
//		first.Next = prev
//		prev = first
//		first = tmp
//	}
//	return prev
//}

// 模拟
func reverseKGroup(head *ListNode, k int) *ListNode {
	// 变量：hair,pre,head,tail
	hair := &ListNode{Next: head} // dummy节点
	pre := hair                   // 记录反转部分的链表的头部的前一个节点
	// head 反转部分的链表的头部，永远在pre的后一个节点上
	// tail 反转部分的链表的尾部
	for head != nil { // 从头部开始
		tail := pre // tail每次循环会重新赋值；重点！
		// 查看剩余部分长度是否大于等于 k
		for i := 0; i < k; i++ {
			tail = tail.Next // 移动tail，向后移动k次，得到需要反转的链表尾部
			if tail == nil {
				return hair.Next
			}
		}

		nex := tail.Next                   // 反转后需要重新连接的下个节点
		head, tail = myReverse(head, tail) // 反转，反转head到tail部分的链表
		// 把子链表重新接回原链表，反转后与原先的链表都已断开
		pre.Next = head
		tail.Next = nex
		// pre和head向后移动k次，都用tail来指定；重点！
		pre = tail       // 指向本次反转链表部分尾部
		head = tail.Next // 下次反转链表部分头部
	}
	return hair.Next
}

// 反转链表
func myReverse(head, tail *ListNode) (*ListNode, *ListNode) {
	prev := tail.Next  // 添加dummy节点，等同于 var prev *ListNode
	curr := head       // 从头开始
	for prev != tail { // tail.Next可能非空！
		next := curr.Next
		// 调换位置
		curr.Next = prev
		prev = curr
		curr = next
	}
	return tail, head // 注意顺序，反转链表后首尾互换
}
